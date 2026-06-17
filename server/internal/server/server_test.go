package server

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

type stubSearcher struct{}

func (stubSearcher) Search(context.Context, string) ([]youtube.Video, error) {
	return []youtube.Video{{ID: "abc123"}}, nil
}

func TestRoutes(t *testing.T) {
	handler := New(
		slog.New(slog.NewTextHandler(io.Discard, nil)),
		stubSearcher{},
		Options{RateLimitRequests: 100, RateLimitWindow: time.Minute},
		nil,
	)

	tests := []struct {
		name       string
		path       string
		wantStatus int
	}{
		{"healthz wired", "/healthz", http.StatusOK},
		{"readyz wired", "/readyz", http.StatusOK},
		{"search wired", "/api/v1/search?q=lofi", http.StatusOK},
		{"search without query is 400", "/api/v1/search", http.StatusBadRequest},
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
