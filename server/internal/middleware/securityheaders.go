package middleware

import "net/http"

// contentSecurityPolicy locks responses to the app's own origin, with a
// deliberate, documented exception for the YouTube IFrame player: it must load
// the player script and be allowed to frame the YouTube origin (PRD §11).
// frame-ancestors 'self' keeps the app itself from being embedded elsewhere.
const contentSecurityPolicy = "default-src 'self'; " +
	"img-src 'self' https://i.ytimg.com data:; " +
	"script-src 'self' https://www.youtube.com https://s.ytimg.com; " +
	"style-src 'self' 'unsafe-inline'; " +
	"frame-src https://www.youtube.com https://www.youtube-nocookie.com; " +
	"connect-src 'self'; " +
	"base-uri 'self'; " +
	"form-action 'self'; " +
	"frame-ancestors 'self'"

// SecurityHeaders sets a baseline of security response headers on every response.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := w.Header()
		h.Set("Content-Security-Policy", contentSecurityPolicy)
		h.Set("X-Content-Type-Options", "nosniff")
		h.Set("X-Frame-Options", "DENY")
		h.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		h.Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}
