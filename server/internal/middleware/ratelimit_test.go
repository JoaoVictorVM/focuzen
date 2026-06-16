package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRateLimitByIP(t *testing.T) {
	limited := RateLimitByIP(2, time.Minute)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	call := func(ip string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodGet, "/api/v1/search?q=x", nil)
		req.RemoteAddr = ip + ":5555"
		rec := httptest.NewRecorder()
		limited.ServeHTTP(rec, req)
		return rec
	}

	assert.Equal(t, http.StatusOK, call("1.2.3.4").Code)
	assert.Equal(t, http.StatusOK, call("1.2.3.4").Code)

	blocked := call("1.2.3.4")
	require.Equal(t, http.StatusTooManyRequests, blocked.Code)

	var body map[string]string
	require.NoError(t, json.NewDecoder(blocked.Body).Decode(&body))
	assert.Equal(t, "rate limit exceeded", body["error"])

	// A different IP has its own budget.
	assert.Equal(t, http.StatusOK, call("5.6.7.8").Code)
}
