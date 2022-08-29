package interval

import (
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"go.uber.org/zap"
)

func todayCountSync() error {
	prCount, err := tidb.QueryTodayPRCount()
	if err != nil {
		logger.Error("query tidb today PR count error", zap.Error(err))
		return err
	}

	developer, err := tidb.QueryTodayDeveloperCount()
	if err != nil {
		logger.Error("query tidb today developer count error", zap.Error(err))
		return err
	}

	err = redis.SetTodayEventCount(prCountAndDeveloperNum2Hash(prCount, developer))
	if err != nil {
		logger.Error("set redis today count error", zap.Error(err))
		return err
	}

	return nil
}

func yearCountSync() error {
	prCount, err := tidb.QueryThisYearPRCount()
	if err != nil {
		logger.Error("query tidb this year PR count error", zap.Error(err))
		return err
	}

	developer, err := tidb.QueryThisYearDeveloperCount()
	if err != nil {
		logger.Error("query tidb this year developer count error", zap.Error(err))
		return err
	}

	err = redis.SetThisYearCount(prCountAndDeveloperNum2Hash(prCount, developer))
	if err != nil {
		logger.Error("set redis this year count error", zap.Error(err))
		return err
	}

	return nil
}

func prCountAndDeveloperNum2Hash(prCount map[string]int, developer int64) map[string]interface{} {
	result := make(map[string]interface{})
	if value, exist := prCount["opened"]; exist {
		result[redis.CountOpenKey] = value
	}

	if value, exist := prCount["closed"]; exist {
		result[redis.CountMergeKey] = value
	}

	result[redis.CountDeveloperKey] = developer
	return result
}
