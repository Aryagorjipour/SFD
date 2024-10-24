package tests

import (
	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	"testing"
)

func TestAddDownload(t *testing.T) {
	dm, err := manager.NewDownloadManager("./test_downloads")
	if err != nil {
		t.Fatalf("Failed to create DownloadManager: %v", err)
	}

	url := "https://example.com/file.zip"
	id := dm.AddDownload(url)
	if id != 1 {
		t.Errorf("Expected ID 1, got %d", id)
	}

	if len(dm.GetTasks()) != 1 {
		t.Errorf("Expected 1 task, got %d", len(dm.GetTasks()))
	}
}
