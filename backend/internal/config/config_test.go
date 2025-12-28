package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Save original env values
	originalEnv := make(map[string]string)
	envVars := []string{
		"APP_ENV",
		"PORT",
		"BLUEPRINT_DB_HOST",
		"BLUEPRINT_DB_PORT",
		"BLUEPRINT_DB_USERNAME",
		"BLUEPRINT_DB_PASSWORD",
		"BLUEPRINT_DB_DATABASE",
		"BLUEPRINT_DB_SCHEMA",
	}

	for _, key := range envVars {
		originalEnv[key] = os.Getenv(key)
		os.Unsetenv(key)
	}
	defer func() {
		// Restore original env values
		for key, value := range originalEnv {
			if value != "" {
				os.Setenv(key, value)
			} else {
				os.Unsetenv(key)
			}
		}
	}()

	tests := []struct {
		name    string
		envVars map[string]string
		wantErr bool
	}{
		{
			name: "valid config with defaults",
			envVars: map[string]string{
				"BLUEPRINT_DB_HOST":     "localhost",
				"BLUEPRINT_DB_USERNAME": "postgres",
				"BLUEPRINT_DB_DATABASE": "mindkeep",
			},
			wantErr: false,
		},
		{
			name: "valid config with all env vars",
			envVars: map[string]string{
				"APP_ENV":               "production",
				"PORT":                  "3000",
				"BLUEPRINT_DB_HOST":     "db.example.com",
				"BLUEPRINT_DB_PORT":     "5432",
				"BLUEPRINT_DB_USERNAME": "user",
				"BLUEPRINT_DB_PASSWORD": "pass",
				"BLUEPRINT_DB_DATABASE": "testdb",
				"BLUEPRINT_DB_SCHEMA":   "public",
			},
			wantErr: false,
		},
		{
			name: "missing required database host",
			envVars: map[string]string{
				"BLUEPRINT_DB_USERNAME": "postgres",
				"BLUEPRINT_DB_DATABASE": "mindkeep",
			},
			wantErr: true,
		},
		{
			name: "invalid port",
			envVars: map[string]string{
				"PORT":                  "invalid",
				"BLUEPRINT_DB_HOST":     "localhost",
				"BLUEPRINT_DB_USERNAME": "postgres",
				"BLUEPRINT_DB_DATABASE": "mindkeep",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			cfg, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && cfg == nil {
				t.Error("Load() returned nil config without error")
				return
			}

			if !tt.wantErr {
				// Verify some values
				if tt.envVars["APP_ENV"] != "" && cfg.AppEnv != tt.envVars["APP_ENV"] {
					t.Errorf("AppEnv = %v, want %v", cfg.AppEnv, tt.envVars["APP_ENV"])
				}
			}

			// Clean up
			for key := range tt.envVars {
				os.Unsetenv(key)
			}
		})
	}
}

func TestIsProduction(t *testing.T) {
	cfg := &Config{AppEnv: "production"}
	if !cfg.IsProduction() {
		t.Error("IsProduction() should return true for production env")
	}

	cfg.AppEnv = "local"
	if cfg.IsProduction() {
		t.Error("IsProduction() should return false for local env")
	}
}

func TestIsLocal(t *testing.T) {
	cfg := &Config{AppEnv: "local"}
	if !cfg.IsLocal() {
		t.Error("IsLocal() should return true for local env")
	}

	cfg.AppEnv = "production"
	if cfg.IsLocal() {
		t.Error("IsLocal() should return false for production env")
	}
}
