package config

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewProducer() *KafkaProducer {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": NewConfig().KAFKA_BROKER_URL,
		"acks":              "all",
	})

	return &KafkaProducer{
		producer: p,
	}
}

func (kp *KafkaProducer) WriteMessage(topic *string, msgInBytes []byte) (*kafka.Message, error) {
	deliverych := make(chan kafka.Event, 10000)

	err := kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value:          msgInBytes,
	}, deliverych)

	e := <-deliverych
	m := e.(*kafka.Message)
	if err != nil {
		log.Println(err)
		return m, err
	}
	close(deliverych)
	return m, nil
}
