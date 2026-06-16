package server

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoutes(t *testing.T) {
	handler := New(slog.New(slog.NewTextHandler(io.Discard, nil)))

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{"healthz wired", "/healthz", http.StatusOK},
		{"readyz wired", "/readyz", http.StatusOK},
		{"unknown route returns 404", "/does-not-exist", http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			require.Equal(t, tt.wantStatus, rec.Code)
		})
	}
}
