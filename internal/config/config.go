package config

import (
	"fmt"
	"os"

	"golang-sandbox/internal/db"
)

// Config holds application configuration.
type Config struct {
	// Server
	Port string

	// Environment
	Env string

	// Database
	Database db.Config
}

// Load loads configuration from environment variables.
func Load() (*Config, error) {
	port := getEnv("PORT", "8080")
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is required")
	}

	// Load database configuration
	dbCfg := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "postgres"),
		DBName:   getEnv("DB_NAME", "golang_sandbox"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	// Validate required database config
	if dbCfg.Host == "" {
		return nil, fmt.Errorf("DB_HOST environment variable is required")
	}

	return &Config{
		Port:     port,
		Env:      getEnv("ENV", "development"),
		Database: dbCfg,
	}, nil
}

// getEnv retrieves an environment variable or returns a fallback value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
