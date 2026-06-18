// Package ui implements the Focuzen terminal UI with Bubble Tea.
package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/JoaoVictorVM/focuzen/cli/internal/audio"
	"github.com/JoaoVictorVM/focuzen/cli/internal/i18n"
)

type tickMsg time.Time

type playbackMsg struct {
	err error
}

// Model is the root Bubble Tea model: the current time, the terminal size used
// to center the view, the radio menu state and the audio player.
type Model struct {
	now           time.Time
	width, height int
	cursor        int
	selected      int
	player        audio.Player
	tr            *i18n.Translator
	err           error
}

// New returns the initial model with no audio selected.
func New(player audio.Player, tr *i18n.Translator) Model {
	return Model{
		now:      time.Now(),
		selected: len(stations) - 1,
		player:   player,
		tr:       tr,
	}
}

// Init starts the clock ticking.
func (m Model) Init() tea.Cmd {
	return tick()
}

// tick schedules the next one-second clock update.
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// playSelected plays (or stops) the selected station off the UI goroutine so a
// slow network request never blocks rendering.
func (m Model) playSelected() tea.Cmd {
	player := m.player
	target := stations[m.selected]
	return func() tea.Msg {
		if target.url == "" {
			player.Stop()
			return playbackMsg{}
		}
		return playbackMsg{err: player.Play(target.url)}
	}
}

// Update advances the clock, tracks the window size and drives the radio menu.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.now = time.Time(msg)
		return m, tick()
	case playbackMsg:
		m.err = msg.err
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(stations)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.cursor
			m.err = nil
			return m, m.playSelected()
		}
	}
	return m, nil
}

// View renders the styled clock and radio menu, centered in the terminal.
func (m Model) View() string {
	sections := []string{
		clockStyle.Render(m.now.Format(m.tr.T("timeFormat"))),
		"",
		m.renderMenu(),
		"",
		hintStyle.Render(m.tr.T("hint")),
	}
	if m.err != nil {
		sections = append(sections, hintStyle.Render(m.tr.T("playbackError")))
	}

	content := lipgloss.JoinVertical(lipgloss.Center, sections...)

	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
