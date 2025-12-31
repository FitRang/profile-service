package eventbus

import (
	"log"

	"github.com/Foxtrot-14/FitRang/profile-service/metrics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	p *kafka.Producer
}

func (e *EventBus) NewProducer() (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": e.cfg.Brokers,
		"security.protocol": "PLAINTEXT",

		"acks":              "all",
		"enable.idempotence": true,
		"retries":           5,
	})

	if err != nil {
		return nil, err
	}

	producer := &Producer{p: p}
	producer.startDeliveryLoop()

	return producer, nil
}

func (p *Producer) startDeliveryLoop() {
	go func() {
		for ev := range p.p.Events() {
			switch e := ev.(type) {

			case *kafka.Message:
				topic := "unknown"
				if e.TopicPartition.Topic != nil {
					topic = *e.TopicPartition.Topic
				}

				if e.TopicPartition.Error != nil {
					log.Printf(
						"kafka delivery failed | topic=%s | error=%v",
						topic,
						e.TopicPartition.Error,
					)

					metrics.KafkaProduceDeliveryTotal.
						WithLabelValues(topic, "failed").
						Inc()

					continue
				}

				metrics.KafkaProduceDeliveryTotal.
					WithLabelValues(topic, "success").
					Inc()
			}
		}
	}()
}

func (p *Producer) Close() {
	log.Println("flushing kafka producer…")
	p.p.Flush(10_000)
	p.p.Close()
}
