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
	"github.com/google/go-github/v45/github"
	"github.com/gorilla/websocket"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"go.uber.org/zap"
	"io"
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

	http.HandleFunc("/sampling", func(w http.ResponseWriter, r *http.Request) {
		samplingHandler(w, r, upgrader)
	})

	http.HandleFunc(readonlyConfig.Server.Health, func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "OK")
	})

	http.HandleFunc(readonlyConfig.Server.SyncEvent, func(w http.ResponseWriter, r *http.Request) {
		if err := syncEvent(); err != nil {
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, "OK")
	})

	port := readonlyConfig.Server.Port
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		logger.Fatal("websocket server start error", zap.Error(err))
	}
	logger.Info("websocket start", zap.Int("port", port))
}

func remain(msg github.Event, eventType, repoName, userName string) bool {
	if len(eventType) > 0 && msg.Type != nil && *msg.Type != eventType {
		return false
	}

	if len(repoName) > 0 && msg.Repo != nil && msg.Repo.Name != nil && *msg.Repo.Name != repoName {
		return false
	}

	if len(userName) > 0 && msg.Actor.Login != nil && *msg.Actor.Login != userName {
		return false
	}

	return true
}
