package eventbus

import (
	"github.com/Foxtrot-14/FitRang/profile-service/metrics"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (p *Producer) Publish(
	topic string,
	key []byte,
	value []byte,
) error {
	err := p.p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}, nil)

	if err != nil {
		metrics.KafkaProduceDeliveryTotal.
			WithLabelValues(topic, "enqueue_error").
			Inc()
		return err
	}

	metrics.KafkaProduceEnqueuedTotal.
		WithLabelValues(topic).
		Inc()

	return nil
}
