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
	redisConnector "github.com/go-redis/redis/v9"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"sync"
)

const (
	eventIDPrefix      = "eid_"
	eventYearPrefix    = "year_"
	repoYearPrefix     = "repo_"
	eventDayPrefix     = "date_"
	distinctPrefix     = "d_"
	eventDailyPrefix   = "daily_year_"
	openPRDailyPrefix  = "daily_open_"
	mergePRDailyPrefix = "daily_merge_"
	closePRDailyPrefix = "daily_close_"
	devDailyPrefix     = "daily_dev_"
	yearSumPrefix      = "year_sum_"
	latestPrefix       = "l_"
)

var (
	client        *redisConnector.Client
	redisInitOnce sync.Once
)

func initClient() {
	redisInitOnce.Do(func() {
		readonlyConfig := config.GetReadonlyConfig()
		client = redisConnector.NewClient(&redisConnector.Options{
			Addr:     readonlyConfig.Redis.Host,
			Password: readonlyConfig.Redis.Password,
			DB:       readonlyConfig.Redis.Db,
		})
	})
}
