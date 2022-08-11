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
    "github.com/gorilla/websocket"
    "github.com/pingcap-inc/ossinsight-plugin/config"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "go.uber.org/zap"
    "net/http"
    "strconv"
)

func createWebsocket() {
    readonlyConfig := config.GetReadonlyConfig()

    upgrader := &websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool { return true },
    }

    http.HandleFunc("/loop", func(w http.ResponseWriter, r *http.Request) {
        loopHandler(w, r, upgrader)
    })

    port := readonlyConfig.Server.Websocket.Port
    err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
    if err != nil {
        logger.Fatal("websocket server start error", zap.Error(err))
    }
    logger.Info("websocket start", zap.Int("port", port))
}
