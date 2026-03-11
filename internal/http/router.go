package http

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"golang-sandbox/internal/db"
)

// Server represents the HTTP server.
type Server struct {
	mux      *http.ServeMux
	logger   *slog.Logger
	database *db.Database
}

// NewServer creates a new HTTP server.
func NewServer(logger *slog.Logger, database *db.Database) *Server {
	mux := http.NewServeMux()
	s := &Server{mux: mux, logger: logger, database: database}

	// Health check endpoints
	mux.HandleFunc("GET /up", s.handleUp)
	mux.HandleFunc("GET /health", s.handleHealth)

	return s
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(w, req)
}

// handleUp handles the simple health check endpoint.
func (s *Server) handleUp(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "up"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		s.logger.Error("failed to encode response", slog.Any("error", err))
	}
}

// handleHealth handles detailed health check with database status.
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"status":   "healthy",
		"database": "healthy",
	}

	// Check database connection
	if s.database != nil {
		sqlDB, err := s.database.DB.DB()
		if err != nil {
			s.logger.Error("failed to get database instance", slog.Any("error", err))
			response["status"] = "unhealthy"
			response["database"] = "unhealthy"
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(response)
			return
		}

		if err := sqlDB.Ping(); err != nil {
			s.logger.Error("database ping failed", slog.Any("error", err))
			response["status"] = "unhealthy"
			response["database"] = "unhealthy"
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		s.logger.Error("failed to encode response", slog.Any("error", err))
	}
}
