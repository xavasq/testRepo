package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB_USER     string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
	DB_PASSWORD string
}

func LoadEnv() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка при загрузке файла .env")
	}
	config := &Config{
		DB_USER:     os.Getenv("DB_USER"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}
	return config
}
