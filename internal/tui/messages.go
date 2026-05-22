package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// tickMsg is sent on a timer to refresh the download list
type tickMsg time.Time

// tickCmd returns a command that waits for a tick
func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// urlsAddedMsg is sent when URLs are successfully added
type urlsAddedMsg struct {
	count int
}

// errorMsg is sent when an error occurs
type errorMsg struct {
	err error
}

// downloadActionMsg is sent when an action is taken on a download
type downloadActionMsg struct {
	downloadID int
	action     string // "pause", "resume", "cancel"
}
