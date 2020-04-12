//The isolation layer can abstract some resource isolation methods.
//Under the premise of correct configuration, different configurations can be obtained through resource isolation methods.
//It is recommended to isolate by cluster....

package isolation

import (
	"errors"
	"fmt"
	"github.com/arcosx/Sirius/config"
	"github.com/arcosx/Sirius/kafka/consumer"
	"github.com/arcosx/Sirius/kafka/producer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// all isolation
var IsolationSet Set

type Isolation struct {
	// isolation key
	keyword       string
	producerProxy *producer.Proxy
	consumerProxy *consumer.Proxy
}

// Set for Isolation
type Set struct {
	isolationMap map[string]*Isolation
}

// functions for IsolationSet
func InitIsolationSet() {
	kafkaConfig := config.C.GetKafkaServiceConfig()
	IsolationSet.isolationMap = make(map[string]*Isolation)
	for _, kfc := range kafkaConfig.Isolations {
		SetIsolation(kfc.Keyword, kfc.ProducerConfig, kfc.ConsumerConfig)
	}

}
func SetIsolation(keyword string, ProducerConfig map[string]interface{}, ConsumerConfig map[string]interface{}) {
	var isolation Isolation
	isolation.keyword = keyword
	isolation.producerProxy = producer.InitProxy(ProducerConfig)
	isolation.consumerProxy = consumer.InitProxy(ConsumerConfig)
	IsolationSet.isolationMap[keyword] = &isolation
}
func (T *Set) GetIsolation(keyword string) (*Isolation, error) {
	if keyword == "" {
		return nil, errors.New("isolation keyword is not a legal value")
	}
	if isolation := T.isolationMap[keyword]; isolation != nil {
		return isolation, nil
	}
	return nil, errors.New(fmt.Sprintf("isolation %s not exist", keyword))
}

func (T *Set) String() string {
	var toStringResult string
	for _, value := range IsolationSet.isolationMap {
		toStringResult += fmt.Sprintf("%s\n", value)
	}
	return toStringResult
}

//functions for Isolation
func (T *Isolation) ProduceAsync(topic string, message []byte) {
	T.producerProxy.ProduceAsync(topic, message)
}

func (T *Isolation) Consume(topic string, ch chan *kafka.Message) {
	go T.consumerProxy.ConsumeTopic(topic, ch)
}

func (T *Isolation) String() string {
	var toStringResult string
	toStringResult = fmt.Sprintf("keyword %s\nproducerProxy %s\nconsumerProxy %s", T.keyword, T.producerProxy, T.consumerProxy)
	return toStringResult
}
