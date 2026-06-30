package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("KEY is not set in .env")
	}

	return &Config{Key: key}, nil
}
