package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	t.Run("hit and miss", func(t *testing.T) {
		c := NewLRU[string, int](2, 0)
		c.Add("a", 1)

		got, ok := c.Get("a")
		require.True(t, ok)
		assert.Equal(t, 1, got)

		_, ok = c.Get("missing")
		assert.False(t, ok)
	})

	t.Run("evicts least-recently-used", func(t *testing.T) {
		c := NewLRU[string, int](2, 0)
		c.Add("a", 1)
		c.Add("b", 2)
		_, _ = c.Get("a") // make "b" the LRU
		c.Add("c", 3)     // should evict "b"

		_, ok := c.Get("b")
		assert.False(t, ok)
		_, ok = c.Get("a")
		assert.True(t, ok)
		_, ok = c.Get("c")
		assert.True(t, ok)
	})

	t.Run("update refreshes value without growing", func(t *testing.T) {
		c := NewLRU[string, int](1, 0)
		c.Add("a", 1)
		c.Add("a", 2)

		got, ok := c.Get("a")
		require.True(t, ok)
		assert.Equal(t, 2, got)
	})

	t.Run("expired entries are a miss", func(t *testing.T) {
		now := time.Unix(0, 0)
		c := NewLRU[string, int](2, time.Minute)
		c.now = func() time.Time { return now }

		c.Add("a", 1)
		now = now.Add(2 * time.Minute)

		_, ok := c.Get("a")
		assert.False(t, ok)
	})
}
