package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

// Config holds all application configuration
type Config struct {
	// Server configuration
	Server ServerConfig

	// Database configuration
	Database DatabaseConfig

	// Application environment
	AppEnv string
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port         int
	IdleTimeout  int // in seconds
	ReadTimeout  int // in seconds
	WriteTimeout int // in seconds
}

// DatabaseConfig holds database connection configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Schema   string
}

// Load reads environment variables and returns a Config struct
func Load() (*Config, error) {
	cfg := &Config{}

	// Load AppEnv
	cfg.AppEnv = getEnv("APP_ENV", "local")

	// Load Server config
	port, err := strconv.Atoi(getEnv("PORT", "8080"))
	if err != nil {
		return nil, fmt.Errorf("invalid PORT: %w", err)
	}
	cfg.Server = ServerConfig{
		Port:         port,
		IdleTimeout:  60, // 1 minute default
		ReadTimeout:  10, // 10 seconds default
		WriteTimeout: 30, // 30 seconds default
	}

	// Load Database config
	cfg.Database = DatabaseConfig{
		Host:     getEnv("BLUEPRINT_DB_HOST", "localhost"),
		Port:     getEnv("BLUEPRINT_DB_PORT", "5432"),
		Username: getEnv("BLUEPRINT_DB_USERNAME", "postgres"),
		Password: getEnv("BLUEPRINT_DB_PASSWORD", ""),
		Database: getEnv("BLUEPRINT_DB_DATABASE", "mindkeep"),
		Schema:   getEnv("BLUEPRINT_DB_SCHEMA", "public"),
	}

	// Validate required fields
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// Validate checks that all required configuration values are present
func (c *Config) Validate() error {
	if c.Server.Port <= 0 {
		return fmt.Errorf("PORT must be a positive integer")
	}

	if c.Database.Host == "" {
		return fmt.Errorf("BLUEPRINT_DB_HOST is required")
	}

	if c.Database.Username == "" {
		return fmt.Errorf("BLUEPRINT_DB_USERNAME is required")
	}

	if c.Database.Database == "" {
		return fmt.Errorf("BLUEPRINT_DB_DATABASE is required")
	}

	return nil
}

// IsProduction returns true if the app is running in production mode
func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}

// IsLocal returns true if the app is running in local/development mode
func (c *Config) IsLocal() bool {
	return c.AppEnv == "local"
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
