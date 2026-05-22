package tui

import (
	"fmt"
	"strings"

	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
)

// View renders the TUI
func (m model) View() string {
	if m.quitting {
		return "Thanks for using SFD!\n"
	}

	var b strings.Builder

	// Render based on input mode
	if m.inputMode == addingURLsMode {
		return m.renderAddURLsView()
	}

	// Render normal view
	b.WriteString(m.renderHeader())
	b.WriteString("\n\n")
	b.WriteString(m.renderDownloads())
	b.WriteString("\n")
	b.WriteString(m.renderFooter())

	return b.String()
}

// renderHeader renders the header section
func (m model) renderHeader() string {
	var b strings.Builder

	title := titleStyle.Render("SFD - Smart File Downloader")
	subtitle := subtitleStyle.Render(fmt.Sprintf("Download Directory: %s", m.manager.DownloadDir))

	b.WriteString(title)
	b.WriteString("\n")
	b.WriteString(subtitle)

	return b.String()
}

// renderDownloads renders the list of downloads
func (m model) renderDownloads() string {
	if len(m.downloads) == 0 {
		return helpStyle.Render("No active downloads. Press 'a' to add URLs.")
	}

	var b strings.Builder

	for i, download := range m.downloads {
		// Cursor indicator
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}

		b.WriteString(cursor)
		b.WriteString(m.renderDownload(download, i == m.cursor))
		b.WriteString("\n\n")
	}

	return b.String()
}

// renderDownload renders a single download item
func (m model) renderDownload(snap manager.DownloadSnapshot, selected bool) string {
	var b strings.Builder

	// Line 1: ID, URL, Status
	displayURL := TrimURLForDisplay(snap.URL)
	if len(displayURL) > 60 {
		displayURL = ShortenURL(displayURL, 60)
	}

	idStr := idStyle.Render(fmt.Sprintf("#%d", snap.ID))
	urlStr := urlStyle.Render(displayURL)
	statusStr := getStatusStyle(snap.Status).Render(snap.Status)

	b.WriteString(fmt.Sprintf("%s %s │ %s", idStr, urlStr, statusStr))
	b.WriteString("\n")

	// Line 2: Progress bar
	b.WriteString(m.renderProgressBar(snap.Progress))
	b.WriteString("\n")

	// Line 3: Action buttons
	b.WriteString(m.renderButtons(snap, selected))

	return b.String()
}

// renderProgressBar renders a progress bar
func (m model) renderProgressBar(progress float64) string {
	width := 50
	filled := int(progress / 100.0 * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	bar := strings.Repeat(progressFullChar, filled)
	empty := strings.Repeat(progressEmptyChar, width-filled)

	progressText := fmt.Sprintf("%.1f%%", progress)

	fullBar := progressBarStyle.Render(bar) + progressEmptyStyle.Render(empty)
	return fmt.Sprintf("%s %s", fullBar, progressText)
}

// renderButtons renders action buttons for a download
func (m model) renderButtons(snap manager.DownloadSnapshot, selected bool) string {
	var buttons []string

	// Determine which buttons are active
	canPause := snap.Status == "Downloading"
	canResume := snap.Status == "Paused"
	canCancel := snap.Status == "Downloading" || snap.Status == "Paused" || snap.Status == "Queued"

	// Render buttons
	pauseBtn := "[P]ause"
	resumeBtn := "[R]esume"
	cancelBtn := "[C]ancel"

	if selected {
		if canPause {
			buttons = append(buttons, activeButtonStyle.Render(pauseBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(pauseBtn))
		}

		if canResume {
			buttons = append(buttons, activeButtonStyle.Render(resumeBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(resumeBtn))
		}

		if canCancel {
			buttons = append(buttons, activeButtonStyle.Render(cancelBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(cancelBtn))
		}
	} else {
		if canPause {
			buttons = append(buttons, buttonStyle.Render(pauseBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(pauseBtn))
		}

		if canResume {
			buttons = append(buttons, buttonStyle.Render(resumeBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(resumeBtn))
		}

		if canCancel {
			buttons = append(buttons, buttonStyle.Render(cancelBtn))
		} else {
			buttons = append(buttons, disabledButtonStyle.Render(cancelBtn))
		}
	}

	return strings.Join(buttons, " ")
}

// renderFooter renders the footer with help text and status messages
func (m model) renderFooter() string {
	var b strings.Builder

	// Show message if any
	if m.message != "" {
		b.WriteString(m.message)
		b.WriteString("\n\n")
	}

	// Help text
	help := helpStyle.Render("Keys: [a]dd URLs • [↑/↓] navigate • [p]ause • [r]esume • [c]ancel • [q]uit")
	b.WriteString(help)

	return footerStyle.Render(b.String())
}

// renderAddURLsView renders the URL input view
func (m model) renderAddURLsView() string {
	var b strings.Builder

	title := titleStyle.Render("Add Batch URLs")
	b.WriteString(title)
	b.WriteString("\n\n")

	instructions := helpStyle.Render("Paste URLs below (one per line). Press Ctrl+S to submit, Esc to cancel.")
	b.WriteString(instructions)
	b.WriteString("\n\n")

	// Render textarea
	b.WriteString(inputFocusedStyle.Render(m.textArea.View()))
	b.WriteString("\n\n")

	// Show URL count
	count := CountURLs(m.textArea.Value())
	countStr := helpStyle.Render(fmt.Sprintf("Valid URLs detected: %d", count))
	b.WriteString(countStr)

	return b.String()
}
