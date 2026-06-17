// Package ui implements the Focuzen terminal UI with Bubble Tea.
package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

// Model is the root Bubble Tea model. For now it only tracks the current time.
type Model struct {
	now time.Time
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

// Update advances the clock and handles quit keys.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		m.now = time.Time(msg)
		return m, tick()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the current clock.
func (m Model) View() string {
	return "\n  " + m.now.Format("15:04:05") + "\n\n  q to quit\n"
}
