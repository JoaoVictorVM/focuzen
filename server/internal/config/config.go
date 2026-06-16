// Package config loads runtime configuration from environment variables.
package config

import (
	"errors"
	"os"
)

// Config holds the runtime configuration loaded from the environment.
type Config struct {
	Port          string
	YouTubeAPIKey string
}

// Load reads configuration from environment variables, applying defaults and
// validating required fields. It fails fast when a required secret is missing.
func Load() (Config, error) {
	cfg := Config{
		Port:          getenv("PORT", "8080"),
		YouTubeAPIKey: os.Getenv("YOUTUBE_API_KEY"),
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
