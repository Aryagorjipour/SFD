package tui_test

import (
	"testing"

	"github.com/Aryagorjipour/smart-file-downloader/internal/tui"
)

func TestTrimURLForDisplay(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "URL with query parameters",
			input:    "https://s1.dlserfes.info/dl14/series/1999/The.Sopranos/S01/720.HardSub/The.Sopranos.S01E13.720p.FilmKio.HardSub.mp4?md5=Ji_BZaxpPWmDOiNJPctXag&u=804110&expires=1779508381",
			expected: "https://s1.dlserfes.info/dl14/series/1999/The.Sopranos/S01/720.HardSub/The.Sopranos.S01E13.720p.FilmKio.HardSub.mp4",
		},
		{
			name:     "URL without query parameters",
			input:    "https://example.com/file.mp4",
			expected: "https://example.com/file.mp4",
		},
		{
			name:     "URL with fragment",
			input:    "https://example.com/file.mp4#section",
			expected: "https://example.com/file.mp4",
		},
		{
			name:     "URL with query and fragment",
			input:    "https://example.com/file.mp4?key=value#section",
			expected: "https://example.com/file.mp4",
		},
		{
			name:     "Invalid URL",
			input:    "not-a-valid-url",
			expected: "not-a-valid-url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tui.TrimURLForDisplay(tt.input)
			if result != tt.expected {
				t.Errorf("TrimURLForDisplay(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTrimPath(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		maxLen   int
		expected string
	}{
		{
			name:     "Path shorter than maxLen",
			path:     "/short/path.txt",
			maxLen:   50,
			expected: "/short/path.txt",
		},
		{
			name:     "Path longer than maxLen",
			path:     "/very/long/path/to/some/file/document.txt",
			maxLen:   25,
			expected: "/very/long...document.txt",
		},
		{
			name:     "Very short maxLen",
			path:     "/path/to/file.txt",
			maxLen:   5,
			expected: "/path",
		},
		{
			name:     "Filename longer than maxLen",
			path:     "/path/verylongfilename.txt",
			maxLen:   15,
			expected: "...filename.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tui.TrimPath(tt.path, tt.maxLen)
			if result != tt.expected {
				t.Errorf("TrimPath(%q, %d) = %q, want %q", tt.path, tt.maxLen, result, tt.expected)
			}
			if len(result) > tt.maxLen {
				t.Errorf("TrimPath result %q exceeds maxLen %d", result, tt.maxLen)
			}
		})
	}
}

func TestExtractFilename(t *testing.T) {
	tests := []struct {
		name     string
		url      string
		expected string
	}{
		{
			name:     "URL with filename",
			url:      "https://example.com/path/to/file.mp4",
			expected: "file.mp4",
		},
		{
			name:     "URL with query parameters",
			url:      "https://example.com/file.mp4?key=value",
			expected: "file.mp4",
		},
		{
			name:     "URL without filename",
			url:      "https://example.com/",
			expected: "download",
		},
		{
			name:     "URL with only domain",
			url:      "https://example.com",
			expected: "download",
		},
		{
			name:     "Invalid URL",
			url:      "not-a-url",
			expected: "not-a-url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tui.ExtractFilename(tt.url)
			if result != tt.expected {
				t.Errorf("ExtractFilename(%q) = %q, want %q", tt.url, result, tt.expected)
			}
		})
	}
}

func TestShortenURL(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		maxLen int
	}{
		{
			name:   "Long URL",
			url:    "https://s1.dlserfes.info/dl14/series/1999/The.Sopranos/S01/720.HardSub/The.Sopranos.S01E13.720p.FilmKio.HardSub.mp4",
			maxLen: 50,
		},
		{
			name:   "Short URL",
			url:    "https://example.com/file.mp4",
			maxLen: 50,
		},
		{
			name:   "Very restrictive maxLen",
			url:    "https://example.com/very/long/path/to/file.mp4",
			maxLen: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tui.ShortenURL(tt.url, tt.maxLen)
			if len(result) > tt.maxLen {
				t.Errorf("ShortenURL result length %d exceeds maxLen %d", len(result), tt.maxLen)
			}
		})
	}
}

func TestCleanURLList(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "List with duplicates",
			input:    []string{"https://example.com/1", "https://example.com/1", "https://example.com/2"},
			expected: []string{"https://example.com/1", "https://example.com/2"},
		},
		{
			name:     "List with empty strings",
			input:    []string{"https://example.com/1", "", "https://example.com/2", "  "},
			expected: []string{"https://example.com/1", "https://example.com/2"},
		},
		{
			name:     "List with whitespace",
			input:    []string{"  https://example.com/1  ", "https://example.com/2"},
			expected: []string{"https://example.com/1", "https://example.com/2"},
		},
		{
			name:     "Empty list",
			input:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tui.CleanURLList(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("CleanURLList length = %d, want %d", len(result), len(tt.expected))
			}
			for i, url := range result {
				if url != tt.expected[i] {
					t.Errorf("CleanURLList[%d] = %q, want %q", i, url, tt.expected[i])
				}
			}
		})
	}
}
