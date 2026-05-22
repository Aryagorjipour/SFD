package tui

import (
	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

// inputMode represents the current input mode of the TUI
type inputMode int

const (
	normalMode inputMode = iota
	addingURLsMode
)

// model is the Bubble Tea model that holds the TUI state
type model struct {
	manager      *manager.DownloadManager
	downloads    []manager.DownloadSnapshot
	cursor       int
	inputMode    inputMode
	textArea     textarea.Model
	message      string
	windowWidth  int
	windowHeight int
	quitting     bool
}

// initialModel creates a new model with the given download manager
func initialModel(mgr *manager.DownloadManager) model {
	ta := textarea.New()
	ta.Placeholder = "Paste URLs here (one per line)..."
	ta.Focus()
	ta.CharLimit = 0
	ta.SetWidth(80)
	ta.SetHeight(10)

	return model{
		manager:   mgr,
		downloads: mgr.Snapshot(),
		cursor:    0,
		inputMode: normalMode,
		textArea:  ta,
		message:   "",
	}
}

// Init is called when the program starts
func (m model) Init() tea.Cmd {
	return tickCmd()
}
