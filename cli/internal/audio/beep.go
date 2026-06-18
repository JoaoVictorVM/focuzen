package audio

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

const (
	sampleRate      = beep.SampleRate(44100)
	resampleQuality = 4
)

// BeepPlayer streams MP3 audio through Beep and the system speaker. Only one
// stream plays at a time; Play replaces whatever is currently playing.
type BeepPlayer struct {
	mu          sync.Mutex
	initialized bool
	current     io.Closer
}

// NewBeepPlayer returns a ready-to-use BeepPlayer.
func NewBeepPlayer() *BeepPlayer {
	return &BeepPlayer{}
}

var _ Player = (*BeepPlayer)(nil)

// Play fetches the stream, decodes it as MP3 and sends it to the speaker,
// resampling to the speaker's rate.
func (p *BeepPlayer) Play(streamURL string) error {
	p.Stop()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, streamURL, http.NoBody)
	if err != nil {
		return fmt.Errorf("audio: build request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("audio: request failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return fmt.Errorf("audio: unexpected status %d", resp.StatusCode)
	}

	streamer, format, err := mp3.Decode(resp.Body)
	if err != nil {
		_ = resp.Body.Close()
		return fmt.Errorf("audio: decode stream: %w", err)
	}

	p.mu.Lock()
	defer p.mu.Unlock()
	if err := p.ensureSpeaker(); err != nil {
		_ = streamer.Close()
		return err
	}

	p.current = streamer
	speaker.Play(beep.Resample(resampleQuality, format.SampleRate, sampleRate, streamer))
	return nil
}

// Stop silences playback and releases the current stream.
func (p *BeepPlayer) Stop() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.initialized {
		speaker.Clear()
	}
	if p.current != nil {
		_ = p.current.Close()
		p.current = nil
	}
}

// ensureSpeaker initializes the speaker once, on first playback.
func (p *BeepPlayer) ensureSpeaker() error {
	if p.initialized {
		return nil
	}
	if err := speaker.Init(sampleRate, sampleRate.N(time.Second/10)); err != nil {
		return fmt.Errorf("audio: init speaker: %w", err)
	}
	p.initialized = true
	return nil
}
