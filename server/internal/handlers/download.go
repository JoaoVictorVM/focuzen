package handlers

import "net/http"

// Download redirects to where the CLI binaries are published (GitHub Releases),
// keeping a stable same-origin link in the frontend.
func Download(url string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	}
}
