package producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type Proxy struct {
	configMap kafka.ConfigMap
	p         *kafka.Producer
}

func InitProxy(configInterfaceMap map[string]interface{}) *Proxy {
	var producerProxy Proxy
	var configMap kafka.ConfigMap
	configMap = make(map[string]kafka.ConfigValue)
	var err error
	for k, v := range configInterfaceMap {
		err = configMap.SetKey(k, v)
		if err != nil {
			log.Error("Init ProducerProxy configInterfaceMap Error", err)
		}
	}
	producerProxy.SetConfigMap(configMap)
	newProducer, err := kafka.NewProducer(&configMap)
	if err != nil {
		log.Error("Init ProducerProxy NewProducer Error", configMap, err)
	}
	producerProxy.SetProducer(newProducer)
	return &producerProxy
}
func (T *Proxy) SetConfigMap(configMap kafka.ConfigMap) {
	T.configMap = configMap
}
func (T *Proxy) SetProducer(p *kafka.Producer) {
	T.p = p
}
func (T *Proxy) String() string {
	var toStringResult string
	for configKey, configValue := range T.configMap {
		temp := fmt.Sprintf("key:%s,value:%s\n", configKey, configValue)
		toStringResult += temp
	}
	return toStringResult
}
func (T *Proxy) ProduceAsync(topic string, message []byte) {
	T.p.ProduceChannel() <- &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message}
}
