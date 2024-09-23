package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

// LoadConfig carrega as variáveis de ambiente
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Erro ao carregar o .env: %v", err)
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/saas_db"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
	}, nil
}

// getEnv busca a variável de ambiente ou retorna o valor padrão
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
