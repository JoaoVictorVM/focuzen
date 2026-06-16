package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealthHandlers(t *testing.T) {
	tests := []struct {
		name            string
		handler         http.HandlerFunc
		wantStatus      int
		wantStatusField string
	}{
		{"healthz reports ok", Healthz, http.StatusOK, "ok"},
		{"readyz reports ready", Readyz, http.StatusOK, "ready"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			tt.handler(rec, req)

			res := rec.Result()
			defer func() { _ = res.Body.Close() }()

			require.Equal(t, tt.wantStatus, res.StatusCode)
			assert.Equal(t, "application/json", res.Header.Get("Content-Type"))

			var body healthResponse
			require.NoError(t, json.NewDecoder(res.Body).Decode(&body))
			assert.Equal(t, tt.wantStatusField, body.Status)
		})
	}
}
