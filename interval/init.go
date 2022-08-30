package interval

import (
    "github.com/pingcap-inc/ossinsight-plugin/config"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
    "time"
)

func InitInterval() {
    intervalConfig := config.GetReadonlyConfig().Interval

    dailyInterval, err := time.ParseDuration(intervalConfig.Daily)
    if err != nil {
        logger.Error("daily interval parse error, use default", zap.Error(err))
        dailyInterval = time.Hour
    }

    go func() {
        for range time.Tick(dailyInterval) {
            retry(dailySync)
        }
    }()
}

func retry(handler func() error) {
    intervalConfig := config.GetReadonlyConfig().Interval
    for i := 0; i < intervalConfig.Retry; i++ {
        if err := handler(); err != nil {
            logger.Error("sync error", zap.Int("round", i), zap.Error(err))
            time.Sleep(time.Duration(intervalConfig.RetryWait) * time.Millisecond)
        } else {
            // success
            break
        }
    }
}
