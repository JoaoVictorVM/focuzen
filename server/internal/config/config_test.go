package config

import (
	"testing"

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
}
