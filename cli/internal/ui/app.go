// Package ui implements the Focuzen terminal UI with Bubble Tea.
package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

// Model is the root Bubble Tea model: the current time, the terminal size used
// to center the view, and the radio menu state.
type Model struct {
	now           time.Time
	width, height int
	cursor        int
	selected      int
}

// New returns the initial model with no audio selected.
func New() Model {
	return Model{now: time.Now(), selected: len(stations) - 1}
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

// Update advances the clock, tracks the window size and drives the radio menu.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.now = time.Time(msg)
		return m, tick()
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
		}
	}
	return m, nil
}

// View renders the styled clock and radio menu, centered in the terminal.
func (m Model) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		clockStyle.Render(m.now.Format("15:04:05")),
		"",
		m.renderMenu(),
		"",
		hintStyle.Render("↑/↓ choose · enter select · q quit"),
	)

	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
