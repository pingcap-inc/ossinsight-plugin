package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"go.uber.org/zap"
)

type FirstResponse struct {
	FirstMessageTag bool         `json:"firstMessageTag"`
	EventList       []tidb.Event `json:"eventList"`
}

func writeFirstResponse(connection *websocket.Conn) error {
	eventList, err := redis.ZSetGetAll()
	if err != nil {
		logger.Error("redis get all error", zap.Error(err))
		return err
	}

	response := FirstResponse{
		FirstMessageTag: true,
		EventList:       eventList,
	}

	payload, _ := json.Marshal(response)

	err = connection.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		logger.Error("write first response", zap.Error(err))
		return err
	}

	return nil
}
