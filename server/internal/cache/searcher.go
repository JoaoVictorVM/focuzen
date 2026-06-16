package cache

import (
	"context"
	"time"

	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

// Searcher decorates a youtube.Searcher with an LRU cache so repeated queries
// are served from memory, protecting the upstream API quota. Failed lookups are
// not cached.
type Searcher struct {
	next  youtube.Searcher
	cache *LRU[string, []youtube.Video]
}

// NewSearcher wraps next with an LRU cache of the given capacity and ttl.
func NewSearcher(next youtube.Searcher, capacity int, ttl time.Duration) *Searcher {
	return &Searcher{
		next:  next,
		cache: NewLRU[string, []youtube.Video](capacity, ttl),
	}
}

var _ youtube.Searcher = (*Searcher)(nil)

// Search returns a cached result when present, otherwise delegates to the
// wrapped searcher and caches a successful response.
func (s *Searcher) Search(ctx context.Context, query string) ([]youtube.Video, error) {
	if cached, ok := s.cache.Get(query); ok {
		return cached, nil
	}

	videos, err := s.next.Search(ctx, query)
	if err != nil {
		return nil, err
	}

	s.cache.Add(query, videos)
	return videos, nil
}
