package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/JoaoVictorVM/focuzen/cli/internal/i18n"
)

// mockPlayer records calls instead of touching real audio.
type mockPlayer struct {
	played []string
	stops  int
}

func (m *mockPlayer) Play(url string) error {
	m.played = append(m.played, url)
	return nil
}

func (m *mockPlayer) Stop() { m.stops++ }

func testModel(player *mockPlayer) Model {
	return New(player, i18n.New())
}

func TestUpdateNavigation(t *testing.T) {
	m := testModel(&mockPlayer{})

	model, _ := m.Update(tea.KeyMsg{Type: tea.KeyDown})
	m = model.(Model)
	assert.Equal(t, 1, m.cursor)

	m.cursor = 0
	model, _ = m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = model.(Model)
	assert.Equal(t, 0, m.cursor, "up at the top stays put")

	for range stations {
		model, _ = m.Update(tea.KeyMsg{Type: tea.KeyDown})
		m = model.(Model)
	}
	assert.Equal(t, len(stations)-1, m.cursor, "down cannot pass the last station")
}

func TestSelectPlaysStation(t *testing.T) {
	player := &mockPlayer{}
	m := testModel(player)
	m.cursor = 0

	model, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = model.(Model)
	require.Equal(t, 0, m.selected)
	require.NotNil(t, cmd)

	_, ok := cmd().(playbackMsg)
	assert.True(t, ok)
	assert.Equal(t, []string{stations[0].url}, player.played)
}

func TestSelectNoAudioStops(t *testing.T) {
	player := &mockPlayer{}
	m := testModel(player)
	m.cursor = len(stations) - 1

	model, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = model.(Model)
	require.Equal(t, len(stations)-1, m.selected)
	require.NotNil(t, cmd)

	cmd()
	assert.Equal(t, 1, player.stops)
	assert.Empty(t, player.played)
}

func TestQuitKeys(t *testing.T) {
	m := testModel(&mockPlayer{})
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	require.NotNil(t, cmd)
	assert.IsType(t, tea.QuitMsg{}, cmd())
}
