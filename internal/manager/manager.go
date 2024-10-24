package manager

import (
	"fmt"
	"github.com/Aryagorjipour/smart-file-downloader/internal/task"
	"sync"
)

type DownloadManager struct {
	Tasks       map[int]*task.DownloadTask
	DownloadDir string
	Mu          sync.Mutex
	NextID      int
}

func NewDownloadManager(downloadDir string) (*DownloadManager, error) {
	// Initialize download directory, create if not exists
	// (Error handling omitted for brevity)
	return &DownloadManager{
		Tasks:       make(map[int]*task.DownloadTask),
		DownloadDir: downloadDir,
		NextID:      1,
	}, nil
}

func (dm *DownloadManager) AddDownload(url string) int {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	id := dm.NextID
	dm.NextID++
	t := task.NewDownloadTask(id, url, dm.DownloadDir)
	dm.Tasks[id] = t
	go t.Start()
	fmt.Printf("Download added with ID %d\n", id)
	return id
}

func (dm *DownloadManager) ListDownloads() {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	fmt.Println("Downloads:")
	for id, t := range dm.Tasks {
		fmt.Printf("ID: %d | URL: %s | Status: %s | Progress: %.2f%%\n", id, t.URL, t.Status, t.Progress)
	}
}

func (dm *DownloadManager) PauseDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Pause()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

func (dm *DownloadManager) ResumeDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Resume()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

func (dm *DownloadManager) CancelDownload(id int) {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	if t, exists := dm.Tasks[id]; exists {
		t.Cancel()
	} else {
		fmt.Printf("Download ID %d not found.\n", id)
	}
}

func (dm *DownloadManager) GetTasks() map[int]*task.DownloadTask {
	dm.Mu.Lock()
	defer dm.Mu.Unlock()
	// Create a copy to prevent race conditions
	copy := make(map[int]*task.DownloadTask)
	for k, v := range dm.Tasks {
		copy[k] = v
	}
	return copy
}
