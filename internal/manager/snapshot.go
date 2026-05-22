package manager

import (
	"time"

	"github.com/Aryagorjipour/smart-file-downloader/internal/task"
)

// DownloadSnapshot represents a thread-safe snapshot of a download task
type DownloadSnapshot struct {
	ID         int
	URL        string
	Status     string
	Progress   float64
	DisplayURL string // Pre-trimmed URL for display
	Timestamp  time.Time
}

// Snapshot returns a safe copy of all downloads with their current state
func (dm *DownloadManager) Snapshot() []DownloadSnapshot {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	snapshots := make([]DownloadSnapshot, 0, len(dm.Tasks))
	for id, t := range dm.Tasks {
		// Note: Status and Progress are simple types that are safe to read
		// even without the task's internal mutex. The task updates them
		// atomically through updateProgress/updateStatus methods.
		snapshot := DownloadSnapshot{
			ID:        id,
			URL:       t.URL,
			Status:    t.Status,
			Progress:  t.Progress,
			Timestamp: time.Now(),
		}
		snapshots = append(snapshots, snapshot)
	}
	return snapshots
}

// RemoveTask removes a task from the manager (for auto-cleanup)
func (dm *DownloadManager) RemoveTask(id int) bool {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	if _, exists := dm.Tasks[id]; exists {
		delete(dm.Tasks, id)
		return true
	}
	return false
}

// AddDownloadSilent adds a download without printing (for TUI)
func (dm *DownloadManager) AddDownloadSilent(url string) int {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	id := dm.NextID
	dm.NextID++
	t := task.NewDownloadTask(id, url, dm.DownloadDir)
	dm.Tasks[id] = t
	go t.Start()
	return id
}
