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
    OpenMap         map[string]string `json:"openMap"`
    MergeMap        map[string]string `json:"mergeMap"`
    DevMap          map[string]string `json:"devMap"`
}

func writeFirstResponse(connection *websocket.Conn) error {
    version := config.GetReadonlyConfig().Api.Version

    eventMap, err := redis.EventNumberGetThisYear()
    if err != nil {
        logger.Error("redis get this year event number error", zap.Error(err))
        return err
    }

    openMap, err := redis.OpenPRNumberGetThisYear()
    if err != nil {
        logger.Error("redis get this year open PR map error", zap.Error(err))
        return err
    }

    mergeMap, err := redis.MergePRNumberGetThisYear()
    if err != nil {
        logger.Error("redis get this year merge map error", zap.Error(err))
        return err
    }

    devMap, err := redis.DeveloperNumberGetThisYear()
    if err != nil {
        logger.Error("redis get this year dev map error", zap.Error(err))
        return err
    }
    response := FirstResponse{
        FirstMessageTag: true,
        APIVersion:      version,
        EventMap:        eventMap,
        OpenMap:         openMap,
        MergeMap:        mergeMap,
        DevMap:          devMap,
    }

    payload, _ := json.Marshal(response)

    err = connection.WriteMessage(websocket.TextMessage, payload)
    if err != nil {
        logger.Error("write first response", zap.Error(err))
        return err
    }

    return nil
}
