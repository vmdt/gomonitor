package config

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewConsumer() *KafkaConsumer {
	config := NewConfig()
	c, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.KAFKA_BROKER_URL,
		"group.id":          config.KAFKA_GROUP_ID,
		"auto.offset.reset": "smallest",
	})

	return &KafkaConsumer{consumer: c}
}

func (kc *KafkaConsumer) ReadMessage(topic string, cb func([]byte)) error {
	run := true
	err := kc.consumer.Subscribe(topic, nil)
	if err != nil {
		return err
	}

	defer kc.consumer.Close()

	for run {
		ev := kc.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			cb(e.Value)
		case *kafka.Error:
			log.Printf("error while consume: %v", e)
			run = false
			return nil
		}
	}

	return nil
}
