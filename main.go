package main

import (
	"github.com/arcosx/Sirius/config"
	kafka "github.com/arcosx/Sirius/kafka/grpc"
	"github.com/arcosx/Sirius/kafka/isolation"
	"github.com/arcosx/Sirius/util"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	util.SetConfig()
	isolation.InitIsolationSet()
}
func main() {
	go kafka.NewKafkaService(config.C.Kafka.Rpc["port"])
	// Temporary wording to prevent program termination
	select {}
}
