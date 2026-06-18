// Command focus runs the Focuzen terminal UI.
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/JoaoVictorVM/focuzen/cli/internal/audio"
	"github.com/JoaoVictorVM/focuzen/cli/internal/i18n"
	"github.com/JoaoVictorVM/focuzen/cli/internal/ui"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run() error {
	player := audio.NewBeepPlayer()
	defer player.Stop()

	_, err := tea.NewProgram(ui.New(player, i18n.New()), tea.WithAltScreen()).Run()
	return err
}
