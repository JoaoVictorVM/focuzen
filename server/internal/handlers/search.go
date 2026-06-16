package handlers

import (
	"log/slog"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/JoaoVictorVM/focuzen/server/internal/youtube"
)

// maxQueryLen bounds the search term, both as input sanitization and to keep
// requests to the upstream API well-formed.
const maxQueryLen = 100

// SearchHandler serves the search endpoint, depending on the Searcher port so
// the upstream YouTube client can be mocked in tests.
type SearchHandler struct {
	searcher youtube.Searcher
}

// NewSearchHandler wires a SearchHandler to a Searcher.
func NewSearchHandler(searcher youtube.Searcher) *SearchHandler {
	return &SearchHandler{searcher: searcher}
}

type searchResults struct {
	Results []youtube.Video `json:"results"`
}

type errorResponse struct {
	Error string `json:"error"`
}

// Search validates the query and proxies it to the Searcher, returning the
// results as JSON. Upstream failures are masked as a generic 502.
func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := strings.TrimSpace(r.URL.Query().Get("q"))
	if query == "" {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "query parameter 'q' is required"})
		return
	}
	if utf8.RuneCountInString(query) > maxQueryLen {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "query parameter 'q' is too long"})
		return
	}

	videos, err := h.searcher.Search(r.Context(), query)
	if err != nil {
		slog.ErrorContext(r.Context(), "search failed", "err", err)
		writeJSON(w, http.StatusBadGateway, errorResponse{Error: "search failed"})
		return
	}

	writeJSON(w, http.StatusOK, searchResults{Results: videos})
}
