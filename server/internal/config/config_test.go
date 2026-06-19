package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	t.Run("applies default port when unset", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("PORT", "")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, "8080", cfg.Port)
		assert.Equal(t, "test-key", cfg.YouTubeAPIKey)
	})

	t.Run("reads port override", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("PORT", "9000")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, "9000", cfg.Port)
	})

	t.Run("fails without youtube api key", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "")

		_, err := Load()
		require.Error(t, err)
	})

	t.Run("applies cache defaults", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("CACHE_SIZE", "")
		t.Setenv("CACHE_TTL", "")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, 256, cfg.CacheSize)
		assert.Equal(t, time.Hour, cfg.CacheTTL)
	})

	t.Run("reads cache overrides", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("CACHE_SIZE", "512")
		t.Setenv("CACHE_TTL", "30m")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Equal(t, 512, cfg.CacheSize)
		assert.Equal(t, 30*time.Minute, cfg.CacheTTL)
	})

	t.Run("applies default download url and allows override", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("DOWNLOAD_URL", "")

		cfg, err := Load()
		require.NoError(t, err)
		assert.Contains(t, cfg.DownloadURL, "github.com/JoaoVictorVM/focuzen/releases")

		t.Setenv("DOWNLOAD_URL", "https://example.com/dl")
		cfg, err = Load()
		require.NoError(t, err)
		assert.Equal(t, "https://example.com/dl", cfg.DownloadURL)
	})

	t.Run("rejects invalid cache size", func(t *testing.T) {
		t.Setenv("YOUTUBE_API_KEY", "test-key")
		t.Setenv("CACHE_SIZE", "not-a-number")

		_, err := Load()
		require.Error(t, err)
	})
}
