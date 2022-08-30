package interval

import (
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "github.com/pingcap-inc/ossinsight-plugin/redis"
    "github.com/pingcap-inc/ossinsight-plugin/tidb"
    "go.uber.org/zap"
)

func dailySync() error {
    events, err := tidb.QueryThisYearDailyEvent()
    if err != nil {
        logger.Error("query tidb this year event error", zap.Error(err))
        return err
    }
    err = redis.EventDailyNumberHSet(events)
    if err != nil {
        logger.Error("set redis this year event error", zap.Error(err))
        return err
    }

    openEvents, mergeEvents, err := tidb.QueryThisYearPRDailyEvent()
    if err != nil {
        logger.Error("query tidb this year pr event error", zap.Error(err))
        return err
    }
    err = redis.OpenPRDailyNumberHSet(openEvents)
    if err != nil {
        logger.Error("set redis this year opened pr event error", zap.Error(err))
        return err
    }
    err = redis.MergePRDailyNumberHSet(mergeEvents)
    if err != nil {
        logger.Error("set redis this year merged pr event error", zap.Error(err))
        return err
    }

    developer, err := tidb.QueryThisYearDeveloperDailyEvent()
    if err != nil {
        logger.Error("query tidb this year developer error", zap.Error(err))
        return err
    }
    developerTotal, err := tidb.QueryThisYearDeveloperCount()
    if err != nil {
        logger.Error("query tidb this year developer count error", zap.Error(err))
        return err
    }
    developer = append(developer, tidb.DailyEvent{
        EventDay: "total",
        Events:   developerTotal,
    })
    err = redis.DeveloperDailyNumberHSet(developer)
    if err != nil {
        logger.Error("set redis this year developer error", zap.Error(err))
        return err
    }

    return nil
}
