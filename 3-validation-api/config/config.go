package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    string
	Password string
	Address  string
}

func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	address := os.Getenv("ADDRESS")

	if email == "" || password == "" || address == "" {
		return nil, fmt.Errorf("EMAIL, PASSWORD, or ADDRESS is not set in .env")
	}

	return &Config{Email: email, Password: password, Address: address}, nil
}
