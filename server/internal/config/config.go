// Package config loads runtime configuration from environment variables.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Config holds the runtime configuration loaded from the environment.
type Config struct {
	Port          string
	YouTubeAPIKey string
	CacheSize     int
	CacheTTL      time.Duration
}

// Load reads configuration from environment variables, applying defaults and
// validating required fields. It fails fast when a required secret is missing.
func Load() (Config, error) {
	cacheSize, err := getenvInt("CACHE_SIZE", 256)
	if err != nil {
		return Config{}, err
	}

	cacheTTL, err := getenvDuration("CACHE_TTL", time.Hour)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Port:          getenv("PORT", "8080"),
		YouTubeAPIKey: os.Getenv("YOUTUBE_API_KEY"),
		CacheSize:     cacheSize,
		CacheTTL:      cacheTTL,
	}

	if cfg.YouTubeAPIKey == "" {
		return Config{}, errors.New("config: YOUTUBE_API_KEY is required")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getenvInt(key string, fallback int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("config: invalid %s: %w", key, err)
	}
	return n, nil
}

func getenvDuration(key string, fallback time.Duration) (time.Duration, error) {
	v := os.Getenv(key)
	if v == "" {
		return fallback, nil
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return 0, fmt.Errorf("config: invalid %s: %w", key, err)
	}
	return d, nil
}
