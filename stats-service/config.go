package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	STATS_PORT string
}

func NewConfig() *Config {
	godotenv.Load()
	return &Config{
		STATS_PORT: os.Getenv("STATS_PORT"),
	}
}
