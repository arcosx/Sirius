//The isolation layer can abstract some resource isolation methods.
//Under the premise of correct configuration, different configurations can be obtained through resource isolation methods.
//It is recommended to isolate by cluster....

package isolation

import (
	"errors"
	"fmt"
	"github.com/arcosx/Sirius/kafka/consumer"
	"github.com/arcosx/Sirius/kafka/producer"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Isolation struct {
	// isolation key
	keyword   string
	configMap map[string]kafka.ConfigMap
	producer  *producer.Producer
	consumer  *consumer.Consumer
}

// Set for Isolation
type Set struct {
	isolationMap map[string]*Isolation
}

func InitIsolationSet(isolationMap map[string]*Isolation) *Set {
	return &Set{isolationMap: isolationMap}
}

func (s *Set) GetIsolation(keyword string) (*Isolation, error) {
	if keyword == "" {
		return nil, errors.New("isolation keyword is not a legal value")
	}
	if isolation := s.isolationMap[keyword]; isolation != nil {
		return isolation, nil
	}
	return nil, errors.New(fmt.Sprintf("isolation %s not exist", keyword))
}

func (i *Isolation) Produce(topic string, message []byte) {

}

func (i *Isolation) Consume(topic string) {
}
