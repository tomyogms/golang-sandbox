package main

import (
	"context"
	"log/slog"
	"os"

	"golang-sandbox/internal/config"
	"golang-sandbox/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	srv := server.New(cfg, logger)

	if err := srv.Start(context.Background()); err != nil {
		logger.Error("server failed", slog.Any("error", err))
		os.Exit(1)
	}
}
