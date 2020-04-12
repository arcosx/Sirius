package consumer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

type Proxy struct {
	configMap kafka.ConfigMap
	c         *kafka.Consumer
}

func InitProxy(configInterfaceMap map[string]interface{}) *Proxy {
	var consumerProxy Proxy
	var configMap kafka.ConfigMap
	configMap = make(map[string]kafka.ConfigValue)
	var err error
	for k, v := range configInterfaceMap {
		err = configMap.SetKey(k, v)
		if err != nil {
			log.Error("Init ConsumerProxy configInterfaceMap Error", err)
		}
	}
	consumerProxy.SetConfigMap(configMap)
	newConsumer, err := kafka.NewConsumer(&configMap)
	if err != nil {
		log.Error("Init ConsumerProxy NewConsumer Error", configMap, err)
	}
	consumerProxy.SetConsumer(newConsumer)
	return &consumerProxy
}

func (T *Proxy) SetConfigMap(configMap kafka.ConfigMap) {
	T.configMap = configMap
}
func (T *Proxy) SetConsumer(c *kafka.Consumer) {
	T.c = c
}
func (T *Proxy) String() string {
	var toStringResult string
	for configKey, configValue := range T.configMap {
		temp := fmt.Sprintf("key:%s,value:%s\n", configKey, configValue)
		toStringResult += temp
	}
	return toStringResult
}

func (T *Proxy) ConsumeTopic(topic string, ch chan *kafka.Message) error {
	err := T.c.Subscribe(topic, nil)
	if err != nil {
		log.Error("ConsumeTopic Subscribe Topic", topic, "Error:", err)
		return err
	}
	for {
		msg, err := T.c.ReadMessage(-1)
		if err == nil {
			log.Info("ConsumeTopic Message on", msg.TopicPartition, string(msg.Value))
			ch <- msg
		} else {
			log.Error("ConsumeTopic ReadMessage Error:", err, msg)
			ch <- msg
			return err
		}
	}
}
