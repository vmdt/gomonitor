package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PRODUCER_PORT    string
	KAFKA_BROKER_URL string
	KAFKA_TOPIC      string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		PRODUCER_PORT:    os.Getenv("PRODUCER_PORT"),
		KAFKA_BROKER_URL: os.Getenv("KAFKA_BROKER_URL"),
		KAFKA_TOPIC:      os.Getenv("KAFKA_TOPIC"),
	}
}
