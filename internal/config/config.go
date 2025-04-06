package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ApiKey    string
	SecretKey string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	apiKey := os.Getenv("DADATA_API_KEY")
	secretKey := os.Getenv("DADATA_SECRET_KEY")

	if apiKey == "" || secretKey == "" {
		log.Fatal("DADATA_API_KEY и DADATA_SECRET_KEY не заданы")
	}

	return &Config{
		ApiKey:    apiKey,
		SecretKey: secretKey,
	}

}
