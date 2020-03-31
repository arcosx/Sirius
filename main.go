package main

import (
	kafka "github.com/arcosx/Sirius/kafka/grpc"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	_ = kafka.NewKafkaServer("8848")
}
