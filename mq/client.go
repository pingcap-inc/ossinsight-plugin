package mq

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"go.uber.org/zap"
	"sync"
)

var (
	client         pulsar.Client
	clientInitOnce sync.Once
	readonlyConfig config.Config
)

// initClient Initial client instance. Use clientInitOnce to ensure initial only once.
func initClient() {
	clientInitOnce.Do(func() {
		readonlyConfig = config.GetReadonlyConfig()

		oauth := pulsar.NewAuthenticationOAuth2(map[string]string{
			"type":       "client_credentials",
			"audience":   readonlyConfig.Pulsar.Audience,
			"privateKey": "file://" + readonlyConfig.Pulsar.Keypath,
		})

		var err error
		client, err = pulsar.NewClient(pulsar.ClientOptions{
			URL:            readonlyConfig.Pulsar.Host,
			Authentication: oauth,
		})
		if err != nil {
			logger.Panic("init pulsar client error", zap.Error(err))
		}
	})
}
