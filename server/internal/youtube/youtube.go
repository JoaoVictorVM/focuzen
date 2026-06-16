// Package youtube provides a small client for the YouTube Data API v3 search
// endpoint, behind an interface so handlers depend on the abstraction and tests
// can mock it.
package youtube

import "context"

// Video is a single search result, reduced to the fields the frontend needs.
type Video struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ChannelTitle string `json:"channelTitle"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

// Searcher looks up videos matching a free-text query. Handlers depend on this
// port; the concrete Client is the production adapter.
type Searcher interface {
	Search(ctx context.Context, query string) ([]Video, error)
}
