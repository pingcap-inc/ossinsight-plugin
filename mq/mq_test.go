package mq

import (
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"go.uber.org/atomic"
	"testing"
	"time"
)

func TestInitClient(t *testing.T) {
	initClient()
}

func TestMessage(t *testing.T) {
	messageNum := 10
	for i := 0; i < messageNum; i++ {
		go func() {
			err := Send([]byte("hello world"), "key", nil)
			if err != nil {
				t.Error(err)
			}
		}()
	}

	receivedNum := atomic.NewInt64(0)
	StartConsume(func(message pulsar.Message) error {
		receivedNum.Add(1)
		fmt.Printf("received %d message\n", receivedNum.Load())
		if string(message.Payload()) != "hello world" || message.Key() != "key" {
			t.Errorf("received msg: %+v\n", message)
		}
		return nil
	})

	time.Sleep(5 * time.Second)
}
