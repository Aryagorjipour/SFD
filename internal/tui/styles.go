package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Color scheme
var (
	// Status colors
	colorDownloading = lipgloss.Color("33")  // Blue
	colorPaused      = lipgloss.Color("214") // Orange
	colorCompleted   = lipgloss.Color("42")  // Green
	colorError       = lipgloss.Color("196") // Red
	colorCanceled    = lipgloss.Color("240") // Gray
	colorQueued      = lipgloss.Color("141") // Purple

	// UI colors
	colorPrimary   = lipgloss.Color("212") // Pink
	colorSecondary = lipgloss.Color("86")  // Cyan
	colorMuted     = lipgloss.Color("240") // Gray
	colorAccent    = lipgloss.Color("205") // Magenta
)

// Styles for different UI elements
var (
	// Header style
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			MarginTop(1).
			MarginBottom(1)

	// Title style
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Underline(true)

	// Subtitle style
	subtitleStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			Italic(true)

	// Help text style
	helpStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	// Button styles
	buttonStyle = lipgloss.NewStyle().
			Foreground(colorAccent).
			Bold(true).
			Padding(0, 1)

	disabledButtonStyle = lipgloss.NewStyle().
				Foreground(colorMuted).
				Padding(0, 1)

	activeButtonStyle = lipgloss.NewStyle().
				Foreground(colorSecondary).
				Bold(true).
				Underline(true).
				Padding(0, 1)

	// Status styles
	statusDownloading = lipgloss.NewStyle().
				Foreground(colorDownloading).
				Bold(true)

	statusPaused = lipgloss.NewStyle().
			Foreground(colorPaused).
			Bold(true)

	statusCompleted = lipgloss.NewStyle().
			Foreground(colorCompleted).
			Bold(true)

	statusError = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	statusCanceled = lipgloss.NewStyle().
			Foreground(colorCanceled)

	statusQueued = lipgloss.NewStyle().
			Foreground(colorQueued)

	// URL style
	urlStyle = lipgloss.NewStyle().
			Foreground(colorSecondary)

	// ID style
	idStyle = lipgloss.NewStyle().
		Foreground(colorMuted).
		Bold(true)

	// Progress bar styles
	progressBarStyle = lipgloss.NewStyle().
				Foreground(colorSecondary)

	progressEmptyStyle = lipgloss.NewStyle().
				Foreground(colorMuted)

	// Input area style
	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Padding(1, 2).
			MarginTop(1)

	inputFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colorSecondary).
				Padding(1, 2).
				MarginTop(1)

	// Footer style
	footerStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			MarginTop(1).
			Padding(1, 0)

	// Error message style
	errorStyle = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	// Success message style
	successStyle = lipgloss.NewStyle().
			Foreground(colorCompleted).
			Bold(true)
)

// Progress bar characters
const (
	progressFullChar  = "█"
	progressEmptyChar = "░"
)

// getStatusStyle returns the appropriate style for a status
func getStatusStyle(status string) lipgloss.Style {
	switch status {
	case "Downloading":
		return statusDownloading
	case "Paused":
		return statusPaused
	case "Completed":
		return statusCompleted
	case "Error":
		return statusError
	case "Canceled":
		return statusCanceled
	case "Queued":
		return statusQueued
	default:
		return lipgloss.NewStyle()
	}
}
