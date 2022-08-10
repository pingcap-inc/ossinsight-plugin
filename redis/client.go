package redis

import (
	"github.com/go-redis/redis/v9"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"sync"
)

const (
	eventIDPrefix = "eid_"
)

var (
	client        *redis.Client
	redisInitOnce sync.Once
)

func initClient() {
	redisInitOnce.Do(func() {
		readonlyConfig := config.GetReadonlyConfig()
		client = redis.NewClient(&redis.Options{
			Addr:     readonlyConfig.Redis.Host,
			Password: readonlyConfig.Redis.Password,
			DB:       readonlyConfig.Redis.Db,
		})
	})
}
