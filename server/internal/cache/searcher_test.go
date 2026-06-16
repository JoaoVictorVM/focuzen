package cache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

type countingSearcher struct {
	calls  int
	result []youtube.Video
	err    error
}

func (s *countingSearcher) Search(context.Context, string) ([]youtube.Video, error) {
	s.calls++
	return s.result, s.err
}

func TestCachingSearcher(t *testing.T) {
	t.Run("serves repeated queries from cache", func(t *testing.T) {
		next := &countingSearcher{result: []youtube.Video{{ID: "abc"}}}
		s := NewSearcher(next, 8, time.Hour)

		first, err := s.Search(context.Background(), "lofi")
		require.NoError(t, err)
		second, err := s.Search(context.Background(), "lofi")
		require.NoError(t, err)

		assert.Equal(t, first, second)
		assert.Equal(t, 1, next.calls, "second identical query should hit the cache")
	})

	t.Run("different queries each reach upstream", func(t *testing.T) {
		next := &countingSearcher{result: []youtube.Video{{ID: "abc"}}}
		s := NewSearcher(next, 8, time.Hour)

		_, _ = s.Search(context.Background(), "lofi")
		_, _ = s.Search(context.Background(), "jazz")

		assert.Equal(t, 2, next.calls)
	})

	t.Run("does not cache failures", func(t *testing.T) {
		next := &countingSearcher{err: errors.New("boom")}
		s := NewSearcher(next, 8, time.Hour)

		_, err := s.Search(context.Background(), "lofi")
		require.Error(t, err)
		_, err = s.Search(context.Background(), "lofi")
		require.Error(t, err)

		assert.Equal(t, 2, next.calls, "failed lookups must not be cached")
	})
}
