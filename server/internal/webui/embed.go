// Package webui embeds the built React SPA and serves it as static files with a
// single-page-app fallback. The dist directory is produced by the Vite build and
// copied here before compiling (it is gitignored) — see the Makefile.
package webui

import (
	"embed"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

//go:embed all:dist
var distFS embed.FS

// Handler returns an http.Handler that serves the embedded SPA: real files when
// they exist, otherwise index.html so client-side routes resolve.
func Handler() (http.Handler, error) {
	root, err := fs.Sub(distFS, "dist")
	if err != nil {
		return nil, err
	}
	return &spaHandler{root: root, files: http.FileServer(http.FS(root))}, nil
}

type spaHandler struct {
	root  fs.FS
	files http.Handler
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
	if name == "" {
		name = "index.html"
	}

	if _, err := fs.Stat(h.root, name); err != nil {
		// Unknown path → serve the SPA entry so client-side routing works.
		r = r.Clone(r.Context())
		r.URL.Path = "/"
	}

	h.files.ServeHTTP(w, r)
}
