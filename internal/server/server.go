package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"golang-sandbox/internal/config"
	"golang-sandbox/internal/db"
	httphandler "golang-sandbox/internal/http"
)

// Server manages the HTTP server lifecycle.
type Server struct {
	config   *config.Config
	logger   *slog.Logger
	server   *http.Server
	database *db.Database
}

// New creates a new server instance.
func New(cfg *config.Config, logger *slog.Logger) *Server {
	// Don't create handler yet, wait for database initialization in Start()
	return &Server{
		config: cfg,
		logger: logger,
		server: &http.Server{
			Addr:         ":" + cfg.Port,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

// Start initializes database connection and starts the server.
func (s *Server) Start(ctx context.Context) error {
	// Initialize database connection
	s.logger.Info("connecting to database", slog.String("host", s.config.Database.Host), slog.String("dbname", s.config.Database.DBName))
	database, err := db.New(s.config.Database)
	if err != nil {
		s.logger.Error("failed to initialize database", slog.Any("error", err))
		return err
	}
	s.database = database

	defer func() {
		if err := s.database.Close(); err != nil {
			s.logger.Error("failed to close database", slog.Any("error", err))
		}
	}()

	// Create HTTP handler with database
	handler := httphandler.NewServer(s.logger, s.database)
	s.server.Handler = handler

	s.logger.Info("starting server", slog.String("addr", s.server.Addr), slog.String("env", s.config.Env))
	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the server and closes database connection.
func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down server")
	shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Close database connection
	if s.database != nil {
		if err := s.database.Close(); err != nil {
			s.logger.Error("failed to close database", slog.Any("error", err))
		}
	}

	return s.server.Shutdown(shutdownCtx)
}

// GetAddr returns the server address.
func (s *Server) GetAddr() string {
	return s.server.Addr
}
