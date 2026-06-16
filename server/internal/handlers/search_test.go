package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

type stubSearcher struct {
	videos []youtube.Video
	err    error
}

func (s stubSearcher) Search(context.Context, string) ([]youtube.Video, error) {
	return s.videos, s.err
}

func TestSearchHandler(t *testing.T) {
	t.Run("returns results for a valid query", func(t *testing.T) {
		want := []youtube.Video{{ID: "abc123", Title: "Lofi Beats", ChannelTitle: "Chillhop", ThumbnailURL: "https://img/abc.jpg"}}
		h := NewSearchHandler(stubSearcher{videos: want})

		rec := doSearch(t, h, "lofi music")

		require.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

		var body searchResults
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
		assert.Equal(t, want, body.Results)
	})

	t.Run("rejects missing or blank query with 400", func(t *testing.T) {
		h := NewSearchHandler(stubSearcher{})

		for _, q := range []string{"", "   "} {
			rec := doSearch(t, h, q)
			require.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("rejects too-long query with 400", func(t *testing.T) {
		h := NewSearchHandler(stubSearcher{})

		rec := doSearch(t, h, strings.Repeat("a", maxQueryLen+1))
		require.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("masks upstream failure as 502", func(t *testing.T) {
		h := NewSearchHandler(stubSearcher{err: errors.New("quota exceeded")})

		rec := doSearch(t, h, "lofi")

		require.Equal(t, http.StatusBadGateway, rec.Code)

		var body errorResponse
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&body))
		assert.Equal(t, "search failed", body.Error)
	})
}

func doSearch(t *testing.T, h *SearchHandler, query string) *httptest.ResponseRecorder {
	t.Helper()

	target := "/api/v1/search?q=" + url.QueryEscape(query)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()

	h.Search(rec, req)
	return rec
}
