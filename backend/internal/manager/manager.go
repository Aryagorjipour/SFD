package manager

import (
	"fmt"
	"github.com/Aryagorjipour/smart-file-downloader/internal/task"
	"sort"
	"sync"
)

type DownloadStatus struct {
	ID       int     `json:"id"`
	URL      string  `json:"url"`
	Progress float64 `json:"progress"`
	Status   string  `json:"status"`
}

// DownloadManager manages download tasks
type DownloadManager struct {
	Tasks       map[int]*task.DownloadTask
	DownloadDir string
	Mu          sync.Mutex
	NextID      int
}

// NewDownloadManager creates a new DownloadManager
func NewDownloadManager(downloadDir string) (*DownloadManager, error) {
	return &DownloadManager{
		Tasks:       make(map[int]*task.DownloadTask),
		DownloadDir: downloadDir,
		NextID:      1,
	}, nil
}

// AddDownload adds a new download task to the manager
func (dm *DownloadManager) AddDownload(url string) int {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	id := dm.NextID
	dm.NextID++
	t := task.NewDownloadTask(id, url, dm.DownloadDir)
	dm.Tasks[id] = t
	go t.Start()
	return id
}

// ListDownloads lists all download tasks
func (dm *DownloadManager) ListDownloads() {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	fmt.Println("Downloads:")
	for id, t := range dm.Tasks {
		fmt.Printf("ID: %d | URL: %s | Status: %s | Progress: %.2f%%\n", id, t.URL, t.Status, t.Progress)
	}
}

// PauseDownload pauses a download task
func (dm *DownloadManager) PauseDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Pause()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

// ResumeDownload resumes a paused download task
func (dm *DownloadManager) ResumeDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Resume()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

// CancelDownload cancels a download task
func (dm *DownloadManager) CancelDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Cancel()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

// GetTasks returns copy of all download tasks for test purposes
func (dm *DownloadManager) GetTasks() map[int]*task.DownloadTask {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	// Create a copyTasks to prevent race conditions
	copyTasks := make(map[int]*task.DownloadTask)
	for k, v := range dm.Tasks {
		copyTasks[k] = v
	}
	return copyTasks
}

// GetAllDownloadStatuses returns the status of all download tasks
func (dm *DownloadManager) GetAllDownloadStatuses() []DownloadStatus {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	var statuses []DownloadStatus
	for id, t := range dm.Tasks {
		status := DownloadStatus{
			ID:       id,
			URL:      t.GetURL(),
			Progress: t.GetProgress(),
			Status:   t.GetStatus(),
		}
		statuses = append(statuses, status)
	}

	sort.Slice(statuses, func(i, j int) bool {
		return statuses[i].ID < statuses[j].ID
	})

	return statuses
}

// ClearCancelledDownloads removes all cancelled download tasks from the manager
func (dm *DownloadManager) ClearCancelledDownloads() {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()

	for id, downloadTask := range dm.Tasks {
		if downloadTask.GetStatus() == "Error" {
			delete(dm.Tasks, id)
		}
	}
}
