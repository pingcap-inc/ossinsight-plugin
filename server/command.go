package main

import (
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"go.uber.org/zap"
)

func syncEvent() error {
	events, err := tidb.QueryThisYearEvent()
	if err != nil {
		logger.Error("query tidb this year event error", zap.Error(err))
		return err
	}

	err = redis.ZSetUpsert(events)
	if err != nil {
		logger.Error("set redis this year event error", zap.Error(err))
		return err
	}

	return nil
}
