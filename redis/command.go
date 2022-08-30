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

func EventNumberHSet(prefix string, events []tidb.DailyEvent) error {
    initClient()

    hashKey := prefix + strconv.Itoa(time.Now().Year())

    eventMap := make(map[string]interface{})
    for _, event := range events {
        eventMap[event.EventDay] = event.Events
    }

    return HSet(hashKey, eventMap)
}

func EventDailyNumberHSet(events []tidb.DailyEvent) error {
    return EventNumberHSet(eventDailyPrefix, events)
}

func OpenPRDailyNumberHSet(events []tidb.DailyEvent) error {
    return EventNumberHSet(openPRDailyPrefix, events)
}

func MergePRDailyNumberHSet(events []tidb.DailyEvent) error {
    return EventNumberHSet(mergePRDailyPrefix, events)
}

func DeveloperDailyNumberHSet(events []tidb.DailyEvent) error {
    return EventNumberHSet(devDailyPrefix, events)
}

func HSet(hashKey string, setMap map[string]interface{}) error {
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

func PRNumberIncrease() error {
    return EventNumberIncrease(eventDailyPrefix)
}

func OpenPRNumberIncrease() error {
    return EventNumberIncrease(openPRDailyPrefix)
}

func MergeNumberIncrease() error {
    return EventNumberIncrease(mergePRDailyPrefix)
}

func DevNumberIncrease() error {
    return EventNumberIncrease(devDailyPrefix)
}

func DevNumberTotalIncrease() error {
    return HIncr(devDailyPrefix, "total")
}

func EventNumberIncrease(prefix string) error {
    return HIncr(prefix, time.Now().Format("2006-01-02"))
}

func HIncr(prefix, key string) error {
    initClient()

    hashKey := prefix + strconv.Itoa(time.Now().Year())
    result := client.HIncrBy(context.Background(), hashKey, key, 1)
    if result.Err() != nil {
        logger.Error("event number increase error", zap.Error(result.Err()))
        return result.Err()
    }

    return nil
}
