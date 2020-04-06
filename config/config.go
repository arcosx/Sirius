package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var C Config

type Config struct {
	Kafka KafkaServiceConfig
}

type KafkaServiceConfig struct {
	Rpc        map[string]string
	Isolations []KafkaIsolation
}

type KafkaIsolation struct {
	Keyword        string
	ProducerConfig map[string]interface{}
	ConsumerConfig map[string]interface{}
}

func InitConfig(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found; ignore error if desired")
		} else {
			log.Fatal("Config file was found but another error was produced")
		}
	}
	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
func (T *Config) GetKafkaServiceConfig() KafkaServiceConfig {
	return T.Kafka
}
