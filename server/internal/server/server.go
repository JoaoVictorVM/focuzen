// Package server wires the chi router, the base middleware stack and routes.
package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"github.com/JoaoVictorVM/focuzen/server/internal/handlers"
	"github.com/JoaoVictorVM/focuzen/server/internal/middleware"
	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

// Options carries the tunables the router needs from configuration.
type Options struct {
	RateLimitRequests int
	RateLimitWindow   time.Duration
}

// New builds the HTTP handler with the base middleware stack and routes.
func New(logger *slog.Logger, searcher youtube.Searcher, opts Options) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.SecurityHeaders)
	r.Use(chimw.RequestID)
	r.Use(requestLogger(logger))
	r.Use(chimw.Recoverer)

	// Health checks stay outside the rate limiter so platform probes are never
	// throttled.
	r.Get("/healthz", handlers.Healthz)
	r.Get("/readyz", handlers.Readyz)

	searchHandler := handlers.NewSearchHandler(searcher)
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.RateLimitByIP(opts.RateLimitRequests, opts.RateLimitWindow))
		r.Get("/search", searchHandler.Search)
	})

	return r
}

// requestLogger emits one structured log line per request, carrying the chi
// request ID so logs correlate with responses.
func requestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			logger.Info("http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.Status(),
				"bytes", ww.BytesWritten(),
				"duration_ms", time.Since(start).Milliseconds(),
				"request_id", chimw.GetReqID(r.Context()),
			)
		})
	}
}
