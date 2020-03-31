package consumer

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	consumer *kafka.Consumer
}
