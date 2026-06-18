package audio

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStopBeforePlayIsSafe(t *testing.T) {
	p := NewBeepPlayer()

	// Stopping before anything played must not panic nor initialize the speaker
	// (so tests never touch real audio hardware).
	p.Stop()

	require.False(t, p.initialized)
	require.Nil(t, p.current)
}
