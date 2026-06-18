// Package ui implements the Focuzen terminal UI with Bubble Tea.
package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

// Model is the root Bubble Tea model. For now it only tracks the current time
// and the terminal size used to center the view.
type Model struct {
	now           time.Time
	width, height int
}

// New returns the initial model.
func New() Model {
	return Model{now: time.Now()}
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

// Update advances the clock, tracks the window size and handles quit keys.
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
		}
	}
	return m, nil
}

// View renders the styled clock, centered in the terminal.
func (m Model) View() string {
	content := lipgloss.JoinVertical(
		lipgloss.Center,
		clockStyle.Render(m.now.Format("15:04:05")),
		"",
		hintStyle.Render("q to quit"),
	)

	if m.width == 0 || m.height == 0 {
		return content
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, content)
}
