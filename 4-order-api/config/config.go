package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}
type AuthConfig struct {
	Secret string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: ", err)
		return nil, err
	}
	dsn := os.Getenv("DB_DSN")
	secret := os.Getenv("AUTH_SECRET")
	if dsn == "" || secret == "" {
		return nil, errors.New("DB_DSN or AUTH_SECRET is not set in .env")
	}
	return &Config{Db: DbConfig{Dsn: dsn}, Auth: AuthConfig{Secret: secret}}, nil
}

func NewConfig() (*Config, error) {
	return LoadConfig()
}
