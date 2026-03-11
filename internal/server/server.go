package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"golang-sandbox/internal/config"
	httphandler "golang-sandbox/internal/http"
)

// Server manages the HTTP server lifecycle.
type Server struct {
	config *config.Config
	logger *slog.Logger
	server *http.Server
}

// New creates a new server instance.
func New(cfg *config.Config, logger *slog.Logger) *Server {
	handler := httphandler.NewServer(logger)

	return &Server{
		config: cfg,
		logger: logger,
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// Start starts the server and blocks until it exits.
func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting server", slog.String("addr", s.server.Addr), slog.String("env", s.config.Env))
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return s.server.Shutdown(shutdownCtx)
}

// GetAddr returns the server address.
func (s *Server) GetAddr() string {
	return s.server.Addr
}
