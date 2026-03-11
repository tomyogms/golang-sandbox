package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Server represents the HTTP server.
type Server struct {
	mux    *http.ServeMux
	logger *slog.Logger
}

// NewServer creates a new HTTP server.
func NewServer(logger *slog.Logger) *Server {
	mux := http.NewServeMux()
	s := &Server{mux: mux, logger: logger}

	// Health check endpoint
	mux.HandleFunc("GET /up", s.handleUp)

	return s
}

// ServeHTTP implements the http.Handler interface.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.mux.ServeHTTP(w, req)
}

// handleUp handles the health check endpoint.
func (s *Server) handleUp(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "up"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		s.logger.Error("failed to encode response", slog.Any("error", err))
	}
}
