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

package tidb

import "github.com/pingcap-inc/ossinsight-plugin/config"

type LanguageEvent struct {
    Language string `json:"language"`
    Events   int    `json:"events"`
}

func QueryTodayLanguageEvent() ([]LanguageEvent, error) {
    initDBOnce.Do(createDB)

    tidbConfig := config.GetReadonlyConfig().Tidb
    rows, err := tidb.Query(tidbConfig.Sql.LanguageToday)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []LanguageEvent
    for rows.Next() {
        event := LanguageEvent{}
        err = rows.Scan(&event.Language, &event.Events)
        if err == nil {
            result = append(result, event)
        } else {
            return result, err
        }
    }

    return result, nil
}
