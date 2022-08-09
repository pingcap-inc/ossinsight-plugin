package mq

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/pingcap-inc/ossinsight-plugin/config"
	"github.com/pingcap-inc/ossinsight-plugin/logger"
	"go.uber.org/zap"
	"sync"
)

var (
	producer         pulsar.Producer
	producerInitOnce sync.Once
)

// initProducer Initial producer instance. Use producerInitOnce to ensure initial only once.
func initProducer() {
	initClient()

	producerInitOnce.Do(func() {
		readonlyConfig := config.GetReadonlyConfig()

		var err error
		producer, err = client.CreateProducer(pulsar.ProducerOptions{
			Topic: readonlyConfig.Pulsar.Producer.Topic,
		})

		if err != nil {
			logger.Panic("init pulsar producer error", zap.Error(err))
		}
	})
}

// Send message send function
func Send(payload []byte, key string, seqID *int64) error {
	initProducer()

	_, err := producer.Send(context.Background(), &pulsar.ProducerMessage{
		Payload:    payload,
		Key:        key,
		SequenceID: seqID,
	})

	if err != nil {
		logger.Error("send message error", zap.Error(err))
		return err
	}

	return nil
}
