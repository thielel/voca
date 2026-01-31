package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port         string
	DatabasePath string
	Environment  string
	OpenAIAPIKey string
}

// Load reads configuration from environment variables
func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "8080"),
		DatabasePath: getEnv("DATABASE_PATH", "./voca.db"),
		Environment:  getEnv("ENVIRONMENT", "development"),
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
