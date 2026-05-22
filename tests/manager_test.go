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

func TestSnapshot(t *testing.T) {
	dm, err := manager.NewDownloadManager("./test_downloads")
	if err != nil {
		t.Fatalf("Failed to create DownloadManager: %v", err)
	}

	// Add multiple downloads
	url1 := "https://example.com/file1.zip"
	url2 := "https://example.com/file2.zip"
	dm.AddDownloadSilent(url1)
	dm.AddDownloadSilent(url2)

	// Get snapshot
	snapshots := dm.Snapshot()

	if len(snapshots) != 2 {
		t.Errorf("Expected 2 snapshots, got %d", len(snapshots))
	}

	// Verify snapshot data
	for _, snap := range snapshots {
		if snap.URL == "" {
			t.Error("Snapshot URL should not be empty")
		}
		if snap.Status == "" {
			t.Error("Snapshot Status should not be empty")
		}
		if snap.Timestamp.IsZero() {
			t.Error("Snapshot Timestamp should not be zero")
		}
	}
}

func TestRemoveTask(t *testing.T) {
	dm, err := manager.NewDownloadManager("./test_downloads")
	if err != nil {
		t.Fatalf("Failed to create DownloadManager: %v", err)
	}

	url := "https://example.com/file.zip"
	id := dm.AddDownloadSilent(url)

	// Verify task exists
	if len(dm.GetTasks()) != 1 {
		t.Errorf("Expected 1 task, got %d", len(dm.GetTasks()))
	}

	// Remove task
	removed := dm.RemoveTask(id)
	if !removed {
		t.Error("Expected RemoveTask to return true")
	}

	// Verify task was removed
	if len(dm.GetTasks()) != 0 {
		t.Errorf("Expected 0 tasks after removal, got %d", len(dm.GetTasks()))
	}

	// Try to remove non-existent task
	removed = dm.RemoveTask(999)
	if removed {
		t.Error("Expected RemoveTask to return false for non-existent task")
	}
}

func TestAddDownloadSilent(t *testing.T) {
	dm, err := manager.NewDownloadManager("./test_downloads")
	if err != nil {
		t.Fatalf("Failed to create DownloadManager: %v", err)
	}

	url := "https://example.com/file.zip"
	id := dm.AddDownloadSilent(url)

	if id != 1 {
		t.Errorf("Expected ID 1, got %d", id)
	}

	if len(dm.GetTasks()) != 1 {
		t.Errorf("Expected 1 task, got %d", len(dm.GetTasks()))
	}

	// Add another to verify ID increment
	id2 := dm.AddDownloadSilent("https://example.com/file2.zip")
	if id2 != 2 {
		t.Errorf("Expected ID 2, got %d", id2)
	}
}
