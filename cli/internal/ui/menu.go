package ui

import "github.com/charmbracelet/lipgloss"

// station is a selectable background sound. The stream URL is added when audio
// playback is wired (Beep), in a later step.
type station struct {
	name string
}

// stations lists the background sounds plus a final "no audio" option.
var stations = []station{
	{name: "Lofi Hip Hop"},
	{name: "Jazz Café"},
	{name: "Rain & Thunder"},
	{name: "No audio"},
}

// renderMenu draws the station list, marking the cursor and the selected item.
func (m Model) renderMenu() string {
	lines := make([]string, len(stations))
	for i, s := range stations {
		cursor := "  "
		if i == m.cursor {
			cursor = cursorStyle.Render("› ")
		}

		name := s.name
		if i == m.selected {
			name = selectedStyle.Render(s.name + " ●")
		}

		lines[i] = cursor + name
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}
