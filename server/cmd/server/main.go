// Command server runs the Focuzen backend HTTP server.
package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JoaoVictorVM/focuzen/server/internal/cache"
	"github.com/JoaoVictorVM/focuzen/server/internal/config"
	"github.com/JoaoVictorVM/focuzen/server/internal/server"
	"github.com/JoaoVictorVM/focuzen/server/internal/webui"
	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

func main() {
	if err := run(); err != nil {
		slog.Error("server exited with error", "err", err)
		os.Exit(1)
	}
}

func run() error {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	addr := ":" + cfg.Port

	searcher := cache.NewSearcher(youtube.NewClient(cfg.YouTubeAPIKey), cfg.CacheSize, cfg.CacheTTL)

	spa, err := webui.Handler()
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr: addr,
		Handler: server.New(logger, searcher, server.Options{
			RateLimitRequests: cfg.RateLimitRequests,
			RateLimitWindow:   cfg.RateLimitWindow,
			DownloadURL:       cfg.DownloadURL,
		}, spa),
		ReadHeaderTimeout: 10 * time.Second,
	}

	serverErr := make(chan error, 1)
	go func() {
		logger.Info("server starting", "addr", addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-serverErr:
		return err
	case <-ctx.Done():
	}

	logger.Info("server shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return srv.Shutdown(shutdownCtx)
}
