package server

import (
	"log/slog"
	"os"
	"testing"

	"golang-sandbox/internal/config"
)

func TestNewServer(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := &config.Config{
		Port: "8080",
		Env:  "development",
	}

	srv := New(cfg, logger)

	if srv == nil {
		t.Error("expected server to be created, got nil")
	}

	if srv.server == nil {
		t.Error("expected http.Server to be created, got nil")
	}

	if srv.GetAddr() != ":8080" {
		t.Errorf("expected addr ':8080', got '%s'", srv.GetAddr())
	}
}

func TestServerConfiguration(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	cfg := &config.Config{
		Port: "3000",
		Env:  "production",
	}

	srv := New(cfg, logger)

	if srv.GetAddr() != ":3000" {
		t.Errorf("expected addr ':3000', got '%s'", srv.GetAddr())
	}
}
