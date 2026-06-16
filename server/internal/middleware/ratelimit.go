package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/httprate"
)

// RateLimitByIP limits requests per client IP over a sliding window, replying
// with a JSON 429 when the limit is exceeded. The client IP is resolved from
// X-Forwarded-For / X-Real-IP (set by the Koyeb proxy) with a fallback to the
// remote address.
func RateLimitByIP(requestLimit int, window time.Duration) func(http.Handler) http.Handler {
	return httprate.Limit(
		requestLimit,
		window,
		httprate.WithKeyByRealIP(),
		httprate.WithLimitHandler(func(w http.ResponseWriter, _ *http.Request) {
			writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded")
		}),
	)
}
