package tui

import (
	"fmt"

	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	tea "github.com/charmbracelet/bubbletea"
)

// Start initializes and runs the TUI
func Start(mgr *manager.DownloadManager) error {
	// Create initial model
	m := initialModel(mgr)

	// Create Bubble Tea program with options
	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),       // Use alternate screen buffer
		tea.WithMouseCellMotion(), // Enable mouse tracking
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running TUI: %w", err)
	}

	return nil
}
