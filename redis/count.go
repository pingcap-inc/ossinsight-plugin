package redis

import (
    "github.com/google/go-github/v47/github"
    "github.com/pingcap-inc/ossinsight-plugin/config"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
    "strconv"
    "time"
)

func AddEventCount(event github.Event) (pr, devDay, devYear, merge, open bool) {
    initClient()

    // All the count metric is for PR
    if event.Type == nil || *event.Type != "PullRequestEvent" {
        return
    }

    payload, err := event.ParsePayload()
    if err != nil {
        logger.Error("parse payload error", zap.Error(err))
        return
    }

    prEvent, ok := payload.(*github.PullRequestEvent)
    if !ok || prEvent == nil {
        logger.Error("pr payload not PullRequestEvent type")
        return
    }

    err = PRNumberIncrease()
    if err != nil {
        logger.Error("increase error", zap.Error(err))
        return
    }

    // This action is a PR
    pr = true

    if event.Actor != nil && event.Actor.ID != nil {
        devDay, devYear = AddDeveloperCount(*event.Actor.ID)
    }

    if prEvent.Action != nil {
        switch *prEvent.Action {
        case "closed":
            if prEvent.PullRequest != nil && prEvent.PullRequest.Merged != nil && *prEvent.PullRequest.Merged {
                MergeNumberIncrease()
                merge = true
            }
        case "opened":
            OpenPRNumberIncrease()
            open = true
        }
    }

    if prEvent.PullRequest != nil && prEvent.PullRequest.Base != nil &&
        prEvent.PullRequest.Base.GetRepo() != nil &&
        prEvent.PullRequest.Base.GetRepo().Language != nil &&
        len(*prEvent.PullRequest.Base.GetRepo().Language) != 0 {
        LanguageTodayIncrease(*prEvent.PullRequest.Base.GetRepo().Language)
        HIncrLatest(*prEvent.PullRequest.Base.GetRepo().Language)
    }

    return
}

// AddDeveloperCount return if this year and today first PR
func AddDeveloperCount(developerID int64) (year, today bool) {
    initClient()

    if exist, err := AddDeveloperThisYear(developerID); err == nil {
        year = !exist
    }
    if exist, err := AddDeveloperToday(developerID); err == nil {
        today = !exist
    }

    return
}

func AddDeveloperThisYear(developerID int64) (bool, error) {
    initClient()

    exist, err := DevelopIDThisYearExists(developerID)
    if err != nil {
        logger.Error("query developer id exist error", zap.Error(err))
        return exist, err
    }

    if !exist {
        if err = DevNumberTotalIncrease(); err != nil {
            logger.Error("hincrby error", zap.Error(err))
            return exist, err
        }
    }

    return exist, nil
}

func AddDeveloperToday(developerID int64) (bool, error) {
    initClient()

    exist, err := DevelopIDTodayExists(developerID)

    if err != nil {
        logger.Error("query developer id exist error", zap.Error(err))
        return exist, err
    }

    if !exist {
        if err = DevNumberIncrease(); err != nil {
            logger.Error("hincrby error", zap.Error(err))
            return exist, err
        }
    }

    return exist, nil
}

func LanguageTodayIncrease(language string) error {
    return HIncrYear(languageTodayPrefix+time.Now().Format("2006-01-02"), language)
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

func DevNumberIncrease() error {
    return EventNumberIncrease(devDailyPrefix)
}

func DevNumberTotalIncrease() error {
    return HIncrYear(devDailyPrefix, "total")
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
