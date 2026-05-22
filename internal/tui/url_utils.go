package tui

import (
	"net/url"
	"path/filepath"
	"strings"
)

// TrimURLForDisplay removes query parameters from URL for cleaner display
func TrimURLForDisplay(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return rawURL // Return original if parsing fails
	}

	// Remove query parameters and fragment
	parsed.RawQuery = ""
	parsed.Fragment = ""

	return parsed.String()
}

// TrimPath truncates a path to maxLen with smart ellipsis placement
func TrimPath(path string, maxLen int) string {
	if len(path) <= maxLen {
		return path
	}

	if maxLen < 10 {
		return path[:maxLen]
	}

	// Extract filename to preserve it
	filename := filepath.Base(path)
	filenameLen := len(filename)

	if filenameLen >= maxLen-3 {
		// If filename itself is too long, truncate it from the end
		return "..." + filename[len(filename)-(maxLen-3):]
	}

	// Show beginning + ... + filename
	// Calculate how much of the path we can show before "..."
	remaining := maxLen - filenameLen - 3 // 3 for "..."
	if remaining < 1 {
		remaining = 1
	}

	return path[:remaining] + "..." + filename
}

// ExtractFilename gets the filename from a URL
func ExtractFilename(rawURL string) string {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "unknown"
	}

	// Get the base name from the path
	filename := filepath.Base(parsed.Path)
	if filename == "" || filename == "." || filename == "/" {
		return "download"
	}

	return filename
}

// ShortenURL creates a shortened version showing domain and filename
func ShortenURL(rawURL string, maxLen int) string {
	if len(rawURL) <= maxLen {
		return rawURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return TrimPath(rawURL, maxLen)
	}

	// Get domain and filename
	domain := parsed.Host
	filename := filepath.Base(parsed.Path)

	// Format: domain/.../filename
	shortened := domain + "/.../" + filename

	if len(shortened) <= maxLen {
		return shortened
	}

	// If still too long, truncate
	return TrimPath(shortened, maxLen)
}

// CleanURLList removes duplicates and empty URLs from a list
func CleanURLList(urls []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, u := range urls {
		u = strings.TrimSpace(u)
		if u != "" && !seen[u] {
			seen[u] = true
			result = append(result, u)
		}
	}

	return result
}
