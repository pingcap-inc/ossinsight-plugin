package redis

import (
    "context"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
    "strconv"
    "time"
)

// EventIDExists judge this event id exists or not
func EventIDExists(id string) (bool, error) {
    initClient()

    doSet, err := client.SetNX(context.Background(), eventIDPrefix+id, "", 12*time.Hour).Result()
    return !doSet, err
    // return HyperLogLogExists(eventIDKey, id)
}

// DevelopIDThisYearExists judge this developer id is exists in this year or not
func DevelopIDThisYearExists(actorID int64) (bool, error) {
    yearKey := eventYearPrefix + strconv.Itoa(time.Now().Year())
    return HyperLogLogExists(yearKey, actorID)
}

// DevelopIDTodayExists judge this developer id is exists today or not
func DevelopIDTodayExists(actorID int64) (bool, error) {
    dateKey := eventDayPrefix + time.Now().Format("2006-01-02")
    return HyperLogLogExists(dateKey, actorID)
}

// HyperLogLogExists check value if exist at this key (use HyperLogLog, may be error)
func HyperLogLogExists(key string, value interface{}) (bool, error) {
    initClient()

    result := client.PFAdd(context.Background(), distinctPrefix+key, value)
    if result.Err() != nil {
        logger.Error("hyperloglog add error", zap.Error(result.Err()))
        return false, result.Err()
    }
    return result.Val() == 0, nil
}
