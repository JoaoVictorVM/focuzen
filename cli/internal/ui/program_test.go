package ui

import (
	"bytes"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/x/exp/teatest"

	"github.com/JoaoVictorVM/focuzen/cli/internal/i18n"
)

func TestProgramRendersAndQuits(t *testing.T) {
	t.Setenv("FOCUZEN_LANG", "en")

	tm := teatest.NewTestModel(t, New(&mockPlayer{}, i18n.New()), teatest.WithInitialTermSize(80, 24))

	teatest.WaitFor(t, tm.Output(), func(b []byte) bool {
		return bytes.Contains(b, []byte("quit"))
	})

	tm.Send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	tm.WaitFinished(t, teatest.WithFinalTimeout(2*time.Second))
}
