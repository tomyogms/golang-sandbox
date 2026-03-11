package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		env         map[string]string
		expectedErr bool
		expectedCfg *Config
	}{
		{
			name: "load with defaults",
			env:  map[string]string{},
			expectedCfg: &Config{
				Port: "8080",
				Env:  "development",
			},
		},
		{
			name: "load with custom port",
			env: map[string]string{
				"PORT": "3000",
			},
			expectedCfg: &Config{
				Port: "3000",
				Env:  "development",
			},
		},
		{
			name: "load with custom env",
			env: map[string]string{
				"ENV": "production",
			},
			expectedCfg: &Config{
				Port: "8080",
				Env:  "production",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()

			for k, v := range tt.env {
				os.Setenv(k, v)
			}

			cfg, err := Load()

			if tt.expectedErr && err == nil {
				t.Errorf("expected error, got nil")
			}

			if !tt.expectedErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if err == nil && cfg != nil {
				if cfg.Port != tt.expectedCfg.Port {
					t.Errorf("expected port %s, got %s", tt.expectedCfg.Port, cfg.Port)
				}
				if cfg.Env != tt.expectedCfg.Env {
					t.Errorf("expected env %s, got %s", tt.expectedCfg.Env, cfg.Env)
				}
			}
		})
	}

	os.Clearenv()
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		fallback string
		setEnv   bool
		envValue string
		expected string
	}{
		{
			name:     "env var exists",
			key:      "TEST_VAR",
			fallback: "default",
			setEnv:   true,
			envValue: "custom",
			expected: "custom",
		},
		{
			name:     "env var doesn't exist",
			key:      "NONEXISTENT_VAR",
			fallback: "default",
			setEnv:   false,
			expected: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()

			if tt.setEnv {
				os.Setenv(tt.key, tt.envValue)
			}

			result := getEnv(tt.key, tt.fallback)

			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}

	os.Clearenv()
}
