// Package audio plays background-sound streams behind a small interface so the
// UI depends on the abstraction and tests can mock playback (no real sound).
package audio

// Player plays a background sound from an MP3 stream URL and stops it.
type Player interface {
	Play(streamURL string) error
	Stop()
}
