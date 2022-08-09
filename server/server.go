package main

import (
	"encoding/json"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pingcap-inc/ossinsight-plugin/fetcher"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"github.com/pingcap-inc/ossinsight-plugin/mq"
	"go.uber.org/zap"
)

func main() {
	go fetcher.InitLoop()
	mq.StartConsume(func(message pulsar.Message) error {
		payload := message.Payload()

		var event fetcher.Event
		err := json.Unmarshal(payload, &event)
		if err != nil {
			logger.Debug("event unmarshal error", zap.Error(err))

			// drop this message
			return nil
		}
		logger.Debug("print message", zap.String("username", event.Actor.Login))

		return nil
	})

	wait := make(chan int)
	<-wait
}
