package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
	"testing"
	"time"
)

var producer Proxy
var testEnvConfig = map[string]interface{}{"bootstrap.servers": "myjd:9092"}
var testTopic = "test"
var message1 = []byte("TestProducer_ProduceAsync " + "time " + strconv.FormatInt(time.Now().UnixNano(), 10))
var message2 = []byte("TestProducer_ProduceAsync " + "time " + strconv.FormatInt(time.Now().UnixNano(), 10))
var message3 = []byte("TestProducer_ProduceAsync " + "time " + strconv.FormatInt(time.Now().UnixNano(), 10))

func TestProducer_ProduceAsync(t *testing.T) {
	producer = *InitProxy(testEnvConfig)
	doneChan := make(chan bool)
	go func() {
		defer close(doneChan)
		for e := range producer.p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					t.Log("Delivery failed", ev.TopicPartition)
				} else {
					t.Log("Delivered message to", ev.TopicPartition)
				}
				return
			default:
				t.Log("Ignore event", ev)
			}
		}
	}()
	producer.ProduceAsync(testTopic, message1)
	producer.ProduceAsync(testTopic, message2)
	producer.ProduceAsync(testTopic, message3)
	_ = <-doneChan

	// Manual check kafka topic,example:
	// ./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test --from-beginning
	// kafkacat -b myjd:9092 -t test
}
