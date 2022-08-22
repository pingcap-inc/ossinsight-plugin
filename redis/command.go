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

package redis

import (
	"context"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/tidb"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// ExistsAndSet judge this id exists or not
// Using `setnx` to discern, and add an eventIDPrefix
func ExistsAndSet(id string) (bool, error) {
	initClient()

	doSet, err := client.SetNX(context.Background(), eventIDPrefix+id, "", 12*time.Hour).Result()
	return !doSet, err
}

func EventNumberHSet(events []tidb.Event) error {
	initClient()

	hashKey := eventYearPrefix + strconv.Itoa(time.Now().Year())

	eventMap := make(map[string]interface{})
	for _, event := range events {
		eventMap[event.EventDay] = event.Events
	}

	result := client.HSet(context.Background(), hashKey, eventMap)
	if result.Err() != nil {
		logger.Error("hash upsert error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func EventNumberGetThisYear() (map[string]string, error) {
	initClient()

	hashKey := eventYearPrefix + strconv.Itoa(time.Now().Year())
	result := client.HGetAll(context.Background(), hashKey)
	if result.Err() != nil {
		logger.Error("get all set error", zap.Error(result.Err()))
		return nil, result.Err()
	}

	return result.Val(), nil
}

func EventNumberIncrease() error {
	initClient()

	hashKey := eventYearPrefix + strconv.Itoa(time.Now().Year())
	result := client.HIncrBy(context.Background(), hashKey, time.Now().Format("2006-01-02"), 1)
	if result.Err() != nil {
		logger.Error("event number increase error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}
