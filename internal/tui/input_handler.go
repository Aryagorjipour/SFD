package tui

import (
	"strings"
)

// ParseBatchURLs parses a multi-line string of URLs and returns a cleaned list
func ParseBatchURLs(input string) []string {
	var urls []string

	// Split by newlines first
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Further split by comma or space
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == ',' || r == ' ' || r == '\t'
		})

		for _, part := range parts {
			part = strings.TrimSpace(part)
			if isValidURL(part) {
				urls = append(urls, part)
			}
		}
	}

	// Clean duplicates and empty strings
	return CleanURLList(urls)
}

// isValidURL checks if a string looks like a URL
func isValidURL(str string) bool {
	str = strings.TrimSpace(str)
	return strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://")
}

// CountURLs counts valid URLs in a string
func CountURLs(input string) int {
	return len(ParseBatchURLs(input))
}
