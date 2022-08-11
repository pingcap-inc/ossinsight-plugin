// Copyright 2022 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "encoding/json"
    "github.com/gorilla/websocket"
    "github.com/pingcap-inc/ossinsight-plugin/fetcher"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

type LoopConfig struct {
    LoopTime  int    `json:"loopTime"`
    EventType string `json:"eventType"`
    RepoName  string `json:"repoName"`
    UserName  string `json:"userName"`
    Detail    bool   `json:"detail"`
}

type LoopResult struct {
    MsgList []fetcher.Event `json:"msgList"`
    TypeMap map[string]int  `json:"typeMap"`
}

func loopHandler(w http.ResponseWriter, r *http.Request, upgrader *websocket.Upgrader) {
    connection, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        logger.Error("upgrade websocket error", zap.Error(err))
        return
    }

    name := strconv.FormatInt(time.Now().UnixNano(), 10) +
        strconv.FormatInt(rand.Int63(), 10)

    configChan := make(chan LoopConfig)
    go readLoopHandler(name, connection, configChan)
    go writeLoopHandler(name, connection, configChan)
}

func writeLoopHandler(name string, connection *websocket.Conn, configChan chan LoopConfig) {
    loopConfig := <-configChan

    listener := make(chan fetcher.Event)
    err := ListenerRegister(name, listener)
    if err != nil {
        logger.Error("listener register error", zap.Error(err))
        return
    }

    var msgList []fetcher.Event
    go func() {
        for {
            msg := <-listener

            // filters
            if len(loopConfig.EventType) > 0 && msg.Type != loopConfig.EventType {
                continue
            }

            if len(loopConfig.RepoName) > 0 && msg.Repo.Name != loopConfig.RepoName {
                continue
            }

            if len(loopConfig.UserName) > 0 && msg.Actor.Login != loopConfig.UserName {
                continue
            }

            msgList = append(msgList, msg)
        }
    }()

    for range time.Tick(time.Duration(loopConfig.LoopTime) * time.Millisecond) {
        typeMap := make(map[string]int)
        for _, msg := range msgList {
            if _, exist := typeMap[msg.Type]; !exist {
                typeMap[msg.Type] = 0
            }

            typeMap[msg.Type] = typeMap[msg.Type] + 1
        }

        result := LoopResult{TypeMap: typeMap}
        if loopConfig.Detail {
            result.MsgList = msgList
        }

        data, err := json.Marshal(result)
        if err != nil {
            logger.Error("marshal error", zap.Error(err))
            return
        }

        msgList = []fetcher.Event{}

        err = connection.WriteMessage(websocket.TextMessage, data)
        if err != nil {
            logger.Error("write error", zap.Error(err))
            return
        }
    }
}

func readLoopHandler(name string, connection *websocket.Conn, configChan chan LoopConfig) {
    defer func() {
        ListenerDelete(name)
        connection.Close()
    }()

    configSet := false
    for {
        msgType, msg, err := connection.ReadMessage()
        if err != nil {
            logger.Error("read message error", zap.Error(err))
            return
        }

        if msgType == websocket.TextMessage {
            if configSet {
                logger.Error("config already set")
                continue
            }
            // Get Config
            logger.Debug("got text message", zap.ByteString("msg", msg))
            loopConfig := new(LoopConfig)
            err = json.Unmarshal(msg, loopConfig)
            if err != nil {
                log.Println("config parse error:", err)
                return
            }

            configSet = true
            configChan <- *loopConfig
        }
    }
}
