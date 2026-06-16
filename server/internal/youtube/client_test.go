package youtube

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const sampleResponse = `{
  "items": [
    {
      "id": { "videoId": "abc123" },
      "snippet": {
        "title": "Lofi Beats",
        "channelTitle": "Chillhop",
        "thumbnails": { "medium": { "url": "https://img/abc.jpg" } }
      }
    },
    {
      "id": { "videoId": "def456" },
      "snippet": {
        "title": "Jazz Cafe",
        "channelTitle": "Cafe Music",
        "thumbnails": { "medium": { "url": "https://img/def.jpg" } }
      }
    }
  ]
}`

func TestClientSearch(t *testing.T) {
	t.Run("maps results and sends expected query params", func(t *testing.T) {
		var gotQuery url.Values
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotQuery = r.URL.Query()
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(sampleResponse))
		}))
		defer srv.Close()

		client := NewClient("test-key", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
		videos, err := client.Search(context.Background(), "lofi music")
		require.NoError(t, err)

		require.Len(t, videos, 2)
		assert.Equal(t, Video{
			ID:           "abc123",
			Title:        "Lofi Beats",
			ChannelTitle: "Chillhop",
			ThumbnailURL: "https://img/abc.jpg",
		}, videos[0])
		assert.Equal(t, "def456", videos[1].ID)

		assert.Equal(t, "test-key", gotQuery.Get("key"))
		assert.Equal(t, "lofi music", gotQuery.Get("q"))
		assert.Equal(t, "video", gotQuery.Get("type"))
		assert.Equal(t, "snippet", gotQuery.Get("part"))
	})

	t.Run("returns error on non-200 status", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		}))
		defer srv.Close()

		client := NewClient("test-key", WithBaseURL(srv.URL), WithHTTPClient(srv.Client()))
		_, err := client.Search(context.Background(), "x")
		require.Error(t, err)
	})
}
