package config

import "os"

type KafkaConfig struct {
	Brokers string
}

func LoadKafkaConfig() KafkaConfig {
	return KafkaConfig{
		Brokers: os.Getenv("KAFKA_URI"),
	}
}
