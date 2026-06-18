package ui

import "github.com/charmbracelet/lipgloss"

// station is a selectable background sound. An empty url means "no audio";
// labelID, when set, is a translatable label instead of a fixed name.
type station struct {
	name    string
	labelID string
	url     string
}

// stations lists the background sounds plus a final "no audio" option. The URLs
// are stable SomaFM MP3 streams as a starting point — swap them freely.
var stations = []station{
	{name: "Lofi Hip Hop", url: "https://ice1.somafm.com/groovesalad-128-mp3"},
	{name: "Jazz Café", url: "https://ice1.somafm.com/secretagent-128-mp3"},
	{name: "Ambient", url: "https://ice1.somafm.com/dronezone-128-mp3"},
	{labelID: "noAudio", url: ""},
}

// renderMenu draws the station list, marking the cursor and the selected item.
func (m Model) renderMenu() string {
	lines := make([]string, len(stations))
	for i, s := range stations {
		cursor := "  "
		if i == m.cursor {
			cursor = cursorStyle.Render("› ")
		}

		label := m.stationLabel(s)
		if i == m.selected {
			label = selectedStyle.Render(label + " ●")
		}

		lines[i] = cursor + label
	}
	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

func (m Model) stationLabel(s station) string {
	if s.labelID != "" {
		return m.tr.T(s.labelID)
	}
	return s.name
}
