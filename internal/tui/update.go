package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles all messages and updates the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyPress(msg)

	case tea.MouseMsg:
		return m.handleMouse(msg)

	case tickMsg:
		return m.handleTick(msg)

	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
		m.windowHeight = msg.Height
		return m, nil

	case urlsAddedMsg:
		m.message = successStyle.Render("✓ Added " + string(rune(msg.count)) + " downloads")
		return m, nil

	case errorMsg:
		m.message = errorStyle.Render("✗ Error: " + msg.err.Error())
		return m, nil
	}

	// Update textarea when in adding URLs mode
	if m.inputMode == addingURLsMode {
		m.textArea, cmd = m.textArea.Update(msg)
		return m, cmd
	}

	return m, nil
}

// handleKeyPress handles keyboard input
func (m model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	// Handle adding URLs mode separately
	if m.inputMode == addingURLsMode {
		switch msg.Type {
		case tea.KeyEsc:
			// Cancel adding URLs
			m.inputMode = normalMode
			m.textArea.Reset()
			m.message = ""
			return m, nil

		case tea.KeyCtrlS:
			// Save/submit URLs
			urls := ParseBatchURLs(m.textArea.Value())
			if len(urls) > 0 {
				for _, url := range urls {
					m.manager.AddDownloadSilent(url)
				}
				m.message = successStyle.Render("✓ Added " + string(rune(len(urls))) + " downloads")
			} else {
				m.message = errorStyle.Render("✗ No valid URLs found")
			}
			m.inputMode = normalMode
			m.textArea.Reset()
			return m, nil
		}

		// Let textarea handle other keys
		var cmd tea.Cmd
		m.textArea, cmd = m.textArea.Update(msg)
		return m, cmd
	}

	// Normal mode key handling
	switch msg.String() {
	case "ctrl+c", "q":
		m.quitting = true
		return m, tea.Quit

	case "a":
		// Enter adding URLs mode
		m.inputMode = addingURLsMode
		m.message = "Paste URLs and press Ctrl+S to submit, Esc to cancel"
		return m, nil

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
		return m, nil

	case "down", "j":
		if m.cursor < len(m.downloads)-1 {
			m.cursor++
		}
		return m, nil

	case "p":
		// Pause selected download
		if m.cursor < len(m.downloads) {
			download := m.downloads[m.cursor]
			if download.Status == "Downloading" {
				m.manager.PauseDownload(download.ID)
				m.message = "Paused download #" + string(rune(download.ID))
			}
		}
		return m, nil

	case "r":
		// Resume selected download
		if m.cursor < len(m.downloads) {
			download := m.downloads[m.cursor]
			if download.Status == "Paused" {
				m.manager.ResumeDownload(download.ID)
				m.message = "Resumed download #" + string(rune(download.ID))
			}
		}
		return m, nil

	case "c":
		// Cancel selected download
		if m.cursor < len(m.downloads) {
			download := m.downloads[m.cursor]
			m.manager.CancelDownload(download.ID)
			m.message = "Canceled download #" + string(rune(download.ID))
		}
		return m, nil
	}

	return m, nil
}

// handleMouse handles mouse events
func (m model) handleMouse(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	if msg.Type != tea.MouseLeft {
		return m, nil
	}

	// Mouse support for clicking on downloads
	// This is a simplified version - more sophisticated click zones
	// could be implemented based on the view layout

	return m, nil
}

// handleTick handles periodic refresh ticks
func (m model) handleTick(msg tickMsg) (tea.Model, tea.Cmd) {
	// Get fresh snapshot
	snapshots := m.manager.Snapshot()

	// Auto-cleanup: remove completed, error, and canceled downloads
	for _, snap := range snapshots {
		if shouldRemove(snap.Status) {
			m.manager.RemoveTask(snap.ID)
		}
	}

	// Get updated snapshot after cleanup
	m.downloads = m.manager.Snapshot()

	// Adjust cursor if needed
	if m.cursor >= len(m.downloads) && len(m.downloads) > 0 {
		m.cursor = len(m.downloads) - 1
	}
	if m.cursor < 0 {
		m.cursor = 0
	}

	return m, tickCmd()
}

// shouldRemove determines if a download should be auto-removed
func shouldRemove(status string) bool {
	return status == "Completed" || status == "Error" || status == "Canceled"
}
