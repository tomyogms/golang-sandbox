package http

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandleUp(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := NewServer(logger)

	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "health check returns up",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "up"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/up", nil)
			w := httptest.NewRecorder()

			server.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			var result map[string]string
			if err := json.NewDecoder(w.Body).Decode(&result); err != nil {
				t.Errorf("failed to decode response: %v", err)
			}

			if result["status"] != tt.expectedBody["status"] {
				t.Errorf("expected status 'up', got '%s'", result["status"])
			}
		})
	}
}

func TestContentType(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	server := NewServer(logger)

	req := httptest.NewRequest("GET", "/up", nil)
	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", contentType)
	}
}
