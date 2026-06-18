package ui

import "github.com/charmbracelet/lipgloss"

// Minimalist palette echoing the web: a warm accent for the clock and a faint
// tone for secondary hints.
var (
	clockStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#d97a45"))
	hintStyle  = lipgloss.NewStyle().Faint(true)
)
