package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultBaseURL  = "https://www.googleapis.com/youtube/v3"
	defaultMaxItems = "10"
	requestTimeout  = 10 * time.Second
)

// Client is the production adapter that talks to the real YouTube Data API v3.
type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// Option customizes a Client. The injection points exist mainly so tests can
// point the client at an httptest server.
type Option func(*Client)

// WithBaseURL overrides the API base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) { c.baseURL = baseURL }
}

// WithHTTPClient overrides the underlying HTTP client.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) { c.httpClient = hc }
}

// NewClient builds a Client with sane defaults.
func NewClient(apiKey string, opts ...Option) *Client {
	c := &Client{
		apiKey:     apiKey,
		baseURL:    defaultBaseURL,
		httpClient: &http.Client{Timeout: requestTimeout},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

var _ Searcher = (*Client)(nil)

// searchResponse mirrors the subset of the search.list payload we consume.
type searchResponse struct {
	Items []struct {
		ID struct {
			VideoID string `json:"videoId"`
		} `json:"id"`
		Snippet struct {
			Title        string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
			Thumbnails   struct {
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
			} `json:"thumbnails"`
		} `json:"snippet"`
	} `json:"items"`
}

// Search performs a search.list call and maps the response to []Video.
func (c *Client) Search(ctx context.Context, query string) ([]Video, error) {
	endpoint, err := url.Parse(c.baseURL + "/search")
	if err != nil {
		return nil, fmt.Errorf("youtube: invalid base url: %w", err)
	}

	q := endpoint.Query()
	q.Set("key", c.apiKey)
	q.Set("part", "snippet")
	q.Set("type", "video")
	q.Set("maxResults", defaultMaxItems)
	q.Set("q", query)
	endpoint.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("youtube: build request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("youtube: request failed: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("youtube: unexpected status %d", resp.StatusCode)
	}

	var payload searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("youtube: decode response: %w", err)
	}

	videos := make([]Video, 0, len(payload.Items))
	for _, item := range payload.Items {
		videos = append(videos, Video{
			ID:           item.ID.VideoID,
			Title:        item.Snippet.Title,
			ChannelTitle: item.Snippet.ChannelTitle,
			ThumbnailURL: item.Snippet.Thumbnails.Medium.URL,
		})
	}
	return videos, nil
}
