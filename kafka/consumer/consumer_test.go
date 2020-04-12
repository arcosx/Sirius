package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"testing"
)

var consumer Proxy
var testEnvConfig = map[string]interface{}{
	"bootstrap.servers": "myjd:9092",
	"group.id":          "myGroup",
}

var testTopic = "test"

func TestProxy_ConsumeTopic(t *testing.T) {
	// exp:bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test
	// > zzz
	consumer = *InitProxy(testEnvConfig)
	readChan := make(chan *kafka.Message)
	go consumer.ConsumeTopic(testTopic, readChan)
	for {
		msg := <-readChan
		t.Log(string(msg.Value))
	}
}
