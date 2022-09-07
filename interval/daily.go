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

	yearSum, err := tidb.QueryThisYearSumCount()
	if err != nil {
		logger.Error("query tidb this year developer count error", zap.Error(err))
		return err
	}

	err = redis.EventNumberHSet(events)
	if err != nil {
		logger.Error("set redis this year event error", zap.Error(err))
		return err
	}

	err = redis.SetYearlyContent(yearSum)
	if err != nil {
		logger.Error("set redis this year sum error", zap.Error(err))
		return err
	}
	return nil
}
