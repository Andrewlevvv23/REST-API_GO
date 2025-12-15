package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	DSN  string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using defaults")
	}

	port := getEnv("PORT", ":8080")
	dsn := getEnv("DSN", "root:root@tcp()")

	return &Config{
		Port: port,
		DSN:  dsn,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
