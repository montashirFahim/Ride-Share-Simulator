package config

import (
	"User/internal/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Cfg *model.Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	Cfg = &model.Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "ride_service"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		AuthUser:   getEnv("BASIC_AUTH_USER", "admin"),
		AuthPass:   getEnv("BASIC_AUTH_PASSWORD", "admin123"),
		RedisHost:  getEnv("REDIS_HOST", "localhost"),
		RedisPort:  getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
