package redis

import (
    "github.com/google/go-github/v45/github"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
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

    if prEvent.Repo != nil && prEvent.Repo.Language != nil && len(*prEvent.Repo.Language) != 0 {
        LanguageTodayIncrease(*prEvent.Repo.Language)
    }
    return
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
        if err = DevNumberTotalIncrease(); err != nil {
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
        if err = DevNumberIncrease(); err != nil {
            logger.Error("hincrby error", zap.Error(err))
            return err
        }
    }

    return nil
}

func LanguageTodayIncrease(language string) error {
    return HIncr(languageTodayPrefix+time.Now().Format("2006-01-02"), language)
}
