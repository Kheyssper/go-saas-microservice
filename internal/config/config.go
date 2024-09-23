package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PostgresHost     string
	PostgresPort     string
	ServerPort       string
}

// LoadConfig loads environment variables
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		PostgresUser:     getEnv("POSTGRES_USER"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD"),
		PostgresDB:       getEnv("POSTGRES_DB"),
		PostgresHost:     getEnv("POSTGRES_HOST"),
		PostgresPort:     getEnv("POSTGRES_PORT"),
		ServerPort:       getEnv("SERVER_PORT"),
	}, nil
}

// getEnv searches for the environment variable or returns an error if not found
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}
