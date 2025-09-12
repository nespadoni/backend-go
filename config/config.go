package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	Database  DatabaseConfig
	JWTSecret string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Load() *Config {
	// Tenta carregar o arquivo .env (ignora erro se não existir)
	if err := godotenv.Load(); err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	return &Config{
		Port: getEnv("PORT", "8080"),
		Database: DatabaseConfig{
			Host:     getEnv("PGHOST", "postgres.railway.internal"),
			Port:     getEnv("PGPORT", "5432"),
			User:     getEnv("PGUSER", "postgres"),
			Password: getEnv("PGPASSWORD", ""),
			DBName:   getEnv("PGDATABASE", "railway"),
		},
		JWTSecret: getEnv("JWT_SECRET", "default_secret_for_dev"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
