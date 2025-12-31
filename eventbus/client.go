package eventbus

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Config struct {
	Brokers  string
}

type EventBus struct {
	cfg   Config
	admin *kafka.AdminClient
}

func NewEventBus(cfg Config) (*EventBus, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
		"security.protocol": "PLAINTEXT",
		"api.version.request": true,
	}

	admin, err := kafka.NewAdminClient(conf)
	if err != nil {
		return nil, err
	}

	return &EventBus{
		cfg:   cfg,
		admin: admin,
	}, nil
}
