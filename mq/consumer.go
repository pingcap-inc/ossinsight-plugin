package mq

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
)

// ConsumeFunc Message deal function
type ConsumeFunc func(pulsar.Message) error

var consumerInitOnce sync.Once

// StartConsume start to consume message. Use consumerInitOnce to ensure initial only once.
// create consume number : readonlyConfig.Pulsar.Consumer.Concurrency
func StartConsume(consumeFunc ConsumeFunc) {
	initClient()

	consumerInitOnce.Do(func() {
		readonlyConfig := config.GetReadonlyConfig()

		for i := 0; i < readonlyConfig.Pulsar.Consumer.Concurrency; i++ {
			go perConsumer(readonlyConfig, consumeFunc)
		}
	})
}

// perConsumer every consumer create function
func perConsumer(readonlyConfig config.Config, consumeFunc ConsumeFunc) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            readonlyConfig.Pulsar.Consumer.Topic,
		SubscriptionName: readonlyConfig.Pulsar.Consumer.Name,
		Type:             pulsar.Shared,
	})

	if err != nil {
		logger.Panic("init pulsar producer error", zap.Error(err))
	}

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			logger.Error("receive message error", zap.Error(err))
		}

		if err := consumeFunc(msg); err != nil {
			logger.Error("consume function error", zap.Error(err))
			consumer.Nack(msg)
		} else {
			consumer.Ack(msg)
		}
	}
}
