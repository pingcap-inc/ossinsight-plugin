package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"go.uber.org/zap"
)

type FirstResponse struct {
	FirstMessageTag bool              `json:"firstMessageTag"`
	APIVersion      int               `json:"apiVersion"`
	EventMap        map[string]string `json:"eventMap"`
	YearCountMap    map[string]string `json:"yearCountMap"`
	DayCountMap     map[string]string `json:"dayCountMap"`
}

func writeFirstResponse(connection *websocket.Conn) error {
	version := config.GetReadonlyConfig().Api.Version

	eventMap, err := redis.EventNumberGetThisYear()
	if err != nil {
		logger.Error("redis get this year event number error", zap.Error(err))
		return err
	}

	yearCountMap, err := redis.GetThisYearEventCount()
	if err != nil {
		logger.Error("redis get this year count map error", zap.Error(err))
		return err
	}

	dayCountMap, err := redis.GetTodayEventCount()
	if err != nil {
		logger.Error("redis get today count map error", zap.Error(err))
		return err
	}

	response := FirstResponse{
		FirstMessageTag: true,
		APIVersion:      version,
		EventMap:        eventMap,
		YearCountMap:    yearCountMap,
		DayCountMap:     dayCountMap,
	}

	payload, _ := json.Marshal(response)

	err = connection.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		logger.Error("write first response", zap.Error(err))
		return err
	}

	return nil
}
