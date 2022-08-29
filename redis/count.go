package redis

import (
	"context"
	"github.com/google/go-github/v45/github"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func SetThisYearCount(hash map[string]interface{}) error {
	return SetCountHash(GetYearCountKey(), hash)
}

func SetTodayEventCount(hash map[string]interface{}) error {
	return SetCountHash(GetDayCountKey(), hash)
}

func GetThisYearEventCount() (map[string]string, error) {
	return HGetAll(GetYearCountKey())
}

func GetTodayEventCount() (map[string]string, error) {
	return HGetAll(GetDayCountKey())
}

func AddEventCount(event github.Event) (pr, devDay, devYear, merge, open bool) {
	initClient()

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

	err = EventNumberIncrease()
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
				AddCountForYearAndDay(CountMergeKey)
				merge = true
			}
		case "opened":
			AddCountForYearAndDay(CountOpenKey)
			open = true
		}
	}

	return
}

func AddCountForYearAndDay(key string) {
	initClient()

	client.HIncrBy(context.Background(), GetYearCountKey(), key, 1)
	client.HIncrBy(context.Background(), GetDayCountKey(), key, 1)
}

func AddDeveloperCount(developerID int64) (year, today bool) {
	initClient()

	if err := AddDeveloperThisYear(developerID); err == nil {
		year = true
	}
	if err := AddDeveloperToday(developerID); err == nil {
		today = true
	}

	return
}

func AddDeveloperThisYear(developerID int64) error {
	initClient()

	exist, err := DevelopIDThisYearExists(developerID)
	if err != nil {
		logger.Error("query developer id exist error", zap.Error(err))
		return err
	}

	if !exist {
		if err = client.HIncrBy(context.Background(), GetYearCountKey(), CountDeveloperKey, 1).Err(); err != nil {
			logger.Error("hincrby error", zap.Error(err))
			return err
		}
	}

	return nil
}

func AddDeveloperToday(developerID int64) error {
	initClient()

	exist, err := DevelopIDTodayExists(developerID)

	if err != nil {
		logger.Error("query developer id exist error", zap.Error(err))
		return err
	}

	if !exist {
		if err = client.HIncrBy(context.Background(), GetDayCountKey(), CountDeveloperKey, 1).Err(); err != nil {
			logger.Error("hincrby error", zap.Error(err))
			return err
		}
	}

	return nil
}

func SetCountHash(key string, hash map[string]interface{}) error {
	initClient()

	result := client.HSet(context.Background(), key, hash)
	if result.Err() != nil {
		logger.Error("hash upsert error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func GetYearCountKey() string {
	return yearCountPrefix + strconv.Itoa(time.Now().Year())
}

func GetDayCountKey() string {
	return dayCountPrefix + time.Now().Format("2006-01-02")
}
