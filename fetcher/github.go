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
    "encoding/json"
    "fmt"
    "github.com/pingcap-inc/ossinsight-plugin/config"
    "github.com/pingcap-inc/ossinsight-plugin/logger"
    "github.com/pingcap-inc/ossinsight-plugin/mq"
    "github.com/pingcap-inc/ossinsight-plugin/redis"
    "go.uber.org/atomic"
    "io/ioutil"
    "net/http"
    "time"

    "go.uber.org/zap"
)

var currentTokenIndex = atomic.NewInt64(0)

func FetchJson(size int) ([]byte, error) {
    req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/events?per_page=%d", size), nil)
    if err != nil {
        logger.Error("create http get request error", zap.Error(err))
        return nil, err
    }

    token, err := getToken()
    if err != nil {
        logger.Error("get token error", zap.Error(err))
        return nil, err
    }

    req.Header.Set("Accept", "application/vnd.github+json")
    req.Header.Set("Authorization", "token "+token)

    timeout := config.GetReadonlyConfig().Github.Loop.Timeout

    client := http.Client{
        Timeout: time.Duration(timeout) * time.Millisecond,
    }

    resp, err := client.Do(req)
    if err != nil {
        logger.Error("get events http request error", zap.Error(err))
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        logger.Error("get events http request response code not 200", zap.Int("code", resp.StatusCode))
        return nil, err
    }

    return body, err
}

func FetchEvents(size int) ([]Event, error) {
    payload, err := FetchJson(size)
    if err != nil {
        logger.Error("fetch json error", zap.Error(err))
        return nil, err
    }

    var events []Event
    err = json.Unmarshal(payload, &events)
    if err != nil {
        logger.Error("events unmarshal error", zap.Error(err))
        return nil, err
    }

    return events, nil
}

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
                exists, err := redis.ExistsAndSet(event.ID)
                if err != nil {
                    logger.Error("redis request error", zap.Error(err))
                    continue
                }

                if exists {
                    logger.Debug("event already exists, skip it", zap.String("id", event.ID))
                    continue
                }

                marshaledEvent, err := json.Marshal(event)
                if err != nil {
                    logger.Error("marshal event error", zap.Error(err))
                    continue
                }

                err = mq.Send(marshaledEvent, "")
                if err != nil {
                    logger.Error("seq id parse error", zap.Error(err))
                    continue
                }
            }
        }()
    }
}

// getToken get config tokens by round-robin, this function is concurrent safe
func getToken() (string, error) {
    tokens := config.GetReadonlyConfig().Github.Tokens
    if len(tokens) == 0 {
        return "", fmt.Errorf("github token empty, please config it")
    }

    return tokens[int(currentTokenIndex.Add(1))%len(tokens)], nil
}
