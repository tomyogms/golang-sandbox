package config

import (
	"fmt"
	"os"
)

// Config holds application configuration.
type Config struct {
	// Server
	Port string

	// Environment
	Env string
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	port := getEnv("PORT", "8080")
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is required")
	}

	return &Config{
		Port: port,
		Env:  getEnv("ENV", "development"),
	}, nil
}

// getEnv retrieves an environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
