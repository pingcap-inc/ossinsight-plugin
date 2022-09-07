package redis

import (
	"github.com/google/go-github/v47/github"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func AddEventCount(event github.Event) map[string]interface{} {
	result := make(map[string]interface{})
	repo, devDay, devYear, additions, deletions := 0, 0, 0, 0, 0

	initClient()

	// All the count metric is for PR
	if event.Type == nil || *event.Type != "PullRequestEvent" {
		return result
	}

	payload, err := event.ParsePayload()
	if err != nil {
		logger.Error("parse payload error", zap.Error(err))
		return result
	}

	prEvent, ok := payload.(*github.PullRequestEvent)
	if !ok || prEvent == nil {
		logger.Error("pr payload not PullRequestEvent type")
		return result
	}

	err = PRNumberIncrease()
	if err != nil {
		logger.Error("increase error", zap.Error(err))
		return result
	}

	// This action is a PR
	result["pr"] = 1

	if event.Actor != nil && event.Actor.ID != nil {
		devDay = AddDeveloperToday(*event.Actor.ID)
		result["devDay"] = devDay
		devYear = DeveloperExistsThisYear(*event.Actor.ID)
		result["devYear"] = devYear
	}

	if prEvent.Action != nil {
		switch *prEvent.Action {
		case "closed":
			if prEvent.PullRequest != nil && prEvent.PullRequest.Merged != nil && *prEvent.PullRequest.Merged {
				MergeNumberIncrease()
				result["merge"] = 1

				if prEvent.PullRequest.Additions != nil && prEvent.PullRequest.Deletions != nil {
					additions = *prEvent.PullRequest.Additions
					deletions = *prEvent.PullRequest.Deletions
				}
			} else {
				CloseNumberIncrease()
				result["close"] = 1
			}
		case "opened":
			OpenPRNumberIncrease()
			result["open"] = 1
		}
	}

	if prEvent.PullRequest != nil && prEvent.PullRequest.Base != nil &&
		prEvent.PullRequest.Base.GetRepo() != nil {
		if prEvent.PullRequest.Base.GetRepo().Language != nil &&
			len(*prEvent.PullRequest.Base.GetRepo().Language) != 0 {
			HIncrLatest(*prEvent.PullRequest.Base.GetRepo().Language)
		}

		if prEvent.PullRequest.Base.GetRepo().ID != nil {
			if repoExist, err := RepoIDThisYearExists(*prEvent.PullRequest.Base.GetRepo().ID); err == nil && !repoExist {
				repo = 1
			}
		}
	}

	result["repoYear"] = repo
	HIncrYearSum(devYear, repo, additions, deletions)
	return result
}

func DeveloperExistsThisYear(developerID int64) int {
	initClient()

	if exist, err := DevelopIDThisYearExists(developerID); err == nil && !exist {
		return 1
	}

	return 0
}

func AddDeveloperToday(developerID int64) int {
	initClient()

	exist, err := DevelopIDTodayExists(developerID)
	if err != nil {
		logger.Error("query developer id exist error", zap.Error(err))
		return 0
	}

	if !exist {
		if err = DevNumberIncrease(); err != nil {
			logger.Error("hincrby error", zap.Error(err))
			return 0
		}
	}

	return 1
}

func PRNumberIncrease() error {
	return EventNumberIncrease(eventDailyPrefix)
}

func OpenPRNumberIncrease() error {
	return EventNumberIncrease(openPRDailyPrefix)
}

func MergeNumberIncrease() error {
	return EventNumberIncrease(mergePRDailyPrefix)
}

func CloseNumberIncrease() error {
	return EventNumberIncrease(closePRDailyPrefix)
}

func DevNumberIncrease() error {
	return EventNumberIncrease(devDailyPrefix)
}

func EventNumberIncrease(prefix string) error {
	return HIncrYear(prefix, time.Now().Format("2006-01-02"))
}

func HIncrYear(prefix, field string) error {
	hashKey := prefix + strconv.Itoa(time.Now().Year())
	return HIncr(hashKey, field)
}

func HIncrLatest(language string) error {
	hashKey := latestPrefix + strconv.FormatInt(time.Now().Unix(), 10)
	logger.Debug("start incr latest language", zap.String("lan", language), zap.String("key", hashKey))
	err := HIncr(hashKey, language)
	if err != nil {
		logger.Debug("incr latest language error", zap.Error(err))
	}
	secondLength := config.GetReadonlyConfig().Interval.LatestDuring
	expireErr := Expire(hashKey, time.Duration(secondLength)*2*time.Second)
	if expireErr != nil {
		logger.Debug("expire language error", zap.Error(expireErr))
	}
	return err
}

func MergeLatestLanguage() (map[string]int, error) {
	secondLength := config.GetReadonlyConfig().Interval.LatestDuring
	currentSecond := time.Now().Unix()

	resultMap, err := MergeScriptRun(latestPrefix, currentSecond-int64(secondLength), currentSecond)
	if err != nil {
		logger.Error("multipart get error", zap.Error(err))
		return resultMap, err
	}

	return resultMap, nil
}

func WatchLanguage() (map[string]string, map[string]string, error) {
	secondLength := config.GetReadonlyConfig().Interval.LatestDuring
	currentSecond := time.Now().Unix()

	deletionsKey := latestPrefix + strconv.FormatInt(currentSecond-int64(secondLength)-1, 10)
	deletionsMap, err := HGetAll(deletionsKey)
	if err != nil {
		logger.Error("get deletions map error", zap.Error(err))
		return nil, nil, err
	}

	additionsKey := latestPrefix + strconv.FormatInt(currentSecond-1, 10)
	additionsMap, err := HGetAll(additionsKey)
	if err != nil {
		logger.Error("get additions map error", zap.Error(err))
		return nil, nil, err
	}

	return deletionsMap, additionsMap, nil
}

func EventNumberHSet(events []tidb.DailyEvent) error {
	if err := eventNumberHSetHandler(eventDailyPrefix, events, func(event tidb.DailyEvent) interface{} {
		merged, _ := strconv.Atoi(event.MergedPRs)
		opened, _ := strconv.Atoi(event.OpenedPRs)
		closed, _ := strconv.Atoi(event.ClosedPRs)
		return merged + opened + closed
	}); err != nil {
		logger.Error("set event daily error", zap.Error(err))
		return err
	}

	if err := eventNumberHSetHandler(openPRDailyPrefix, events, func(event tidb.DailyEvent) interface{} {
		return event.OpenedPRs
	}); err != nil {
		logger.Error("set opened pr event daily error", zap.Error(err))
		return err
	}

	if err := eventNumberHSetHandler(mergePRDailyPrefix, events, func(event tidb.DailyEvent) interface{} {
		return event.MergedPRs
	}); err != nil {
		logger.Error("set merged pr event daily error", zap.Error(err))
		return err
	}

	if err := eventNumberHSetHandler(closePRDailyPrefix, events, func(event tidb.DailyEvent) interface{} {
		return event.ClosedPRs
	}); err != nil {
		logger.Error("set closed pr event daily error", zap.Error(err))
		return err
	}

	if err := eventNumberHSetHandler(devDailyPrefix, events, func(event tidb.DailyEvent) interface{} {
		return event.Developers
	}); err != nil {
		logger.Error("set developer daily error", zap.Error(err))
		return err
	}

	return nil
}

func eventNumberHSetHandler(prefix string, events []tidb.DailyEvent,
	parseFunc func(tidb.DailyEvent) interface{}) error {

	hashKey := prefix + strconv.Itoa(time.Now().Year())

	eventMap := make(map[string]interface{})
	for _, event := range events {
		eventMap[event.EventDay] = parseFunc(event)
	}

	return HSet(hashKey, eventMap)
}

func SetYearlyContent(content tidb.YearlyContent) error {
	hashKey := yearSumPrefix + strconv.Itoa(time.Now().Year())
	return HSet(hashKey, map[string]interface{}{
		"dev":       content.Developers,
		"repo":      content.Repos,
		"additions": content.Additions,
		"deletions": content.Deletions,
	})
}

func HIncrYearSum(dev, repo, additions, deletions int) {
	hashKey := yearSumPrefix + strconv.Itoa(time.Now().Year())
	if dev != 0 {
		HIncr(hashKey, "dev")
	}

	if repo != 0 {
		HIncr(hashKey, "repo")
	}

	if additions != 0 {
		HIncrBy(hashKey, "additions", int64(additions))
	}

	if deletions != 0 {
		HIncrBy(hashKey, "deletions", int64(deletions))
	}
}

func HGetAllYearSum() (map[string]string, error) {
	hashKey := yearSumPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}
