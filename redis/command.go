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
	redisConnector "github.com/go-redis/redis/v9"
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

func ZSetUpsert(events []tidb.Event) error {
	initClient()

	zSetKey := eventYearPrefix + strconv.Itoa(time.Now().Year())

	members := make([]redisConnector.Z, len(events), len(events))
	for i, event := range events {
		members[i] = redisConnector.Z{
			Score:  float64(event.Events),
			Member: event.EventDay,
		}
	}

	result := client.ZAdd(context.Background(), zSetKey, members...)
	if result.Err() != nil {
		logger.Error("sorted set upsert error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func ZSetGetAll() ([]tidb.Event, error) {
	initClient()

	zSetKey := eventYearPrefix + strconv.Itoa(time.Now().Year())
	result := client.ZRangeWithScores(context.Background(), zSetKey, 0, -1)
	if result.Err() != nil {
		logger.Error("get all sorted set with score error", zap.Error(result.Err()))
		return nil, result.Err()
	}

	members := result.Val()
	events := make([]tidb.Event, len(members), len(members))
	for i, member := range members {
		if stringMember, memberIsString := member.Member.(string); memberIsString {
			events[i] = tidb.Event{
				EventDay: stringMember,
				Events:   int(member.Score),
			}
		}
	}

	return events, nil
}

func ZSetIncrease() error {
	initClient()

	zSetKey := eventYearPrefix + strconv.Itoa(time.Now().Year())
	result := client.ZIncrBy(context.Background(), zSetKey, 1, time.Now().Format("2006-01-02"))
	if result.Err() != nil {
		logger.Error("sorted set member increase score error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}
