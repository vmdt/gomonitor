package handlers

import (
	"encoding/json"
	"log"
	"producer/config"
)

var kafkaProducer = config.NewProducer()

func PushDataToKafka(data any) error {
	statsInBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("error while json marshal: %v", err)
		return err
	}

	m, err := kafkaProducer.WriteMessage(&config.NewConfig().KAFKA_TOPIC, statsInBytes)
	if err != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return err
	} else {
		log.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	return nil
}
