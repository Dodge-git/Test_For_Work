package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		DBHost: os.Getenv("DB_HOST"),
		DBUser: os.Getenv("DB_USER"),
		DBPass: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
		DBPort: os.Getenv("DB_PORT"),
	}

	if cfg.DBHost == "" {
		log.Fatal("DB_HOST not set")
	}

	return cfg
}