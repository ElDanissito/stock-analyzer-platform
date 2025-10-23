package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	Environment    string
	DatabaseURL    string
	StockAPIURL    string
	StockAPIKey    string
	AllowedOrigins string
}

func Load() *Config {
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return &Config{
		Port:           getEnv("PORT", "8080"),
		Environment:    getEnv("ENV", "development"),
		DatabaseURL:    getEnv("DATABASE_URL", ""),
		StockAPIURL:    getEnv("STOCK_API_URL", ""),
		StockAPIKey:    getEnv("STOCK_API_KEY", ""),
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "*"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
