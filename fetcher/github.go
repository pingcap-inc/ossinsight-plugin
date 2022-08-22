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

package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v45/github"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/mq"
	"github.com/pingcap-inc/ossinsight-plugin/redis"
	"go.uber.org/atomic"
	"golang.org/x/oauth2"
	"sync"
	"time"

	"go.uber.org/zap"
)

var (
	// currentClientIndex round-robin index for client
	currentClientIndex = atomic.NewInt64(0)
	// githubClientList create client by each token
	githubClientList []*github.Client
	// initClientOnce create client once
	initClientOnce sync.Once
)

// InitLoop create loop
func InitLoop() {
	readonlyConfig := config.GetReadonlyConfig()
	loopBreak := readonlyConfig.Github.Loop.Break
	for range time.Tick(time.Duration(loopBreak) * time.Millisecond) {
		go func() {
			logger.Debug("start to request github", zap.Int("loopBreak", loopBreak))

			events, err := FetchEvents(100)
			if err != nil {
				logger.Error("fetch event error", zap.Error(err))
				return
			}

			for _, event := range events {
				if event.ID == nil {
					logger.Error("event id not exist")
					continue
				}

				exists, err := redis.ExistsAndSet(*event.ID)
				if err != nil {
					logger.Error("redis request error", zap.Error(err))
					continue
				}

				if exists {
					logger.Debug("event already exists, skip it", zap.String("id", *event.ID))
					continue
				}

				// add calculator number
				err = redis.EventNumberIncrease()
				if err != nil {
					logger.Error("redis request error", zap.Error(err))
					// continue for send message, calculator doesn't matter
				}

				marshaledEvent, err := json.Marshal(event)
				if err != nil {
					logger.Error("marshal event error", zap.Error(err))
					continue
				}

				err = mq.Send(marshaledEvent, "")
				if err != nil {
					logger.Error("send message error", zap.Error(err))
					continue
				}
			}
		}()
	}
}

// FetchEvents fetch GitHub events
func FetchEvents(perPage int) ([]*github.Event, error) {
	client, err := getClient()
	if err != nil {
		logger.Error("get github client error", zap.Error(err))
		return nil, err
	}

	if client == nil {
		logger.Error("github client is nil")
		return nil, fmt.Errorf("github client is nil")
	}

	events, _, err := client.Activity.ListEvents(
		context.Background(), &github.ListOptions{PerPage: perPage})

	if err != nil {
		logger.Error("request github events API error", zap.Error(err))
		return nil, err
	}

	return events, nil
}

// getClient get client by round-robin, this function is concurrent safe
func getClient() (*github.Client, error) {
	initClientOnce.Do(createClients)

	currentCallNum := int(currentClientIndex.Add(1))
	return githubClientList[currentCallNum%len(githubClientList)], nil
}

func createClients() {
	tokens := config.GetReadonlyConfig().Github.Tokens

	for i := range tokens {
		staticTokenSource := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: tokens[i]},
		)
		tc := oauth2.NewClient(context.Background(), staticTokenSource)

		githubClientList = append(githubClientList, github.NewClient(tc))
	}
}
