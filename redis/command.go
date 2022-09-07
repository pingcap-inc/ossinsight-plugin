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
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"go.uber.org/zap"
	"strconv"
	"time"
)

func HSet(hashKey string, setMap map[string]interface{}) error {
	initClient()

	result := client.HSet(context.Background(), hashKey, setMap)
	if result.Err() != nil {
		logger.Error("hash upsert error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func EventNumberGetThisYear() (map[string]string, error) {
	hashKey := eventDailyPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}

func OpenPRNumberGetThisYear() (map[string]string, error) {
	hashKey := openPRDailyPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}

func MergePRNumberGetThisYear() (map[string]string, error) {
	hashKey := mergePRDailyPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}

func ClosePRNumberGetThisYear() (map[string]string, error) {
	hashKey := closePRDailyPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}

func DeveloperNumberGetThisYear() (map[string]string, error) {
	hashKey := devDailyPrefix + strconv.Itoa(time.Now().Year())
	return HGetAll(hashKey)
}

func HGetAll(key string) (map[string]string, error) {
	initClient()

	result := client.HGetAll(context.Background(), key)
	if result.Err() != nil {
		logger.Error("get all set error", zap.Error(result.Err()))
		return nil, result.Err()
	}

	return result.Val(), nil
}

func HIncr(key, field string) error {
	return HIncrBy(key, field, 1)
}

func HIncrBy(key, field string, incr int64) error {
	initClient()

	result := client.HIncrBy(context.Background(), key, field, 1)
	if result.Err() != nil {
		logger.Error("hash value increase error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func Expire(key string, expiration time.Duration) error {
	initClient()

	result := client.Expire(context.Background(), key, expiration)
	if result.Err() != nil {
		logger.Error("set expire error", zap.Error(result.Err()))
		return result.Err()
	}

	return nil
}

func MergeScriptRun(prefix string, start, end int64) (map[string]int, error) {
	initClient()

	mergeLatest := redisConnector.NewScript(config.GetReadonlyConfig().Redis.Lua.MergeLatest)
	result, err := mergeLatest.Run(context.Background(), client,
		[]string{prefix, strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)}).Slice()

	if err != nil {
		logger.Error("script run get error", zap.Error(err))
		return nil, err
	}

	resultMap := make(map[string]int)
	for i := range result {
		if i%2 != 0 {
			language, ok := result[i-1].(string)
			if !ok {
				logger.Error("language not a string", zap.Any("item", result[i-1]))
				continue
			}

			num, ok := result[i].(int64)
			if !ok {
				logger.Error("appear number not an int64", zap.Any("item", result[i]))
				continue
			}

			if _, exist := resultMap[language]; !exist {
				resultMap[language] = 0
			}

			resultMap[language] = resultMap[language] + int(num)
		}
	}

	return resultMap, err
}
