package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"go.uber.org/zap"
)

type FirstResponse struct {
	FirstMessageTag bool              `json:"firstMessageTag"`
	EventMap        map[string]string `json:"eventMap"`
}

func writeFirstResponse(connection *websocket.Conn) error {
	eventMap, err := redis.EventNumberGetThisYear()
	if err != nil {
		logger.Error("redis get this year event number error", zap.Error(err))
		return err
	}

	response := FirstResponse{
		FirstMessageTag: true,
		EventMap:        eventMap,
	}

	payload, _ := json.Marshal(response)

	err = connection.WriteMessage(websocket.TextMessage, payload)
	if err != nil {
		logger.Error("write first response", zap.Error(err))
		return err
	}

	return nil
}
