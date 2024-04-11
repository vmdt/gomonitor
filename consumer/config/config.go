package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	CONSUMER_PORT    string
	KAFKA_BROKER_URL string
	KAFKA_GROUP_ID   string
	KAFKA_TOPIC      string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		CONSUMER_PORT:    os.Getenv("CONSUMER_PORT"),
		KAFKA_BROKER_URL: os.Getenv("KAFKA_BROKER_URL"),
		KAFKA_GROUP_ID:   os.Getenv("KAFKA_GROUP_ID"),
		KAFKA_TOPIC:      os.Getenv("KAFKA_TOPIC"),
	}
}
