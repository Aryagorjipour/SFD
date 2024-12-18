package manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/Aryagorjipour/smart-file-downloader/internal/task"
)

// DownloadManager manages download tasks
type DownloadManager struct {
	Tasks       map[int]*task.DownloadTask
	DownloadDir string
	Mu          sync.Mutex
	NextID      int
}

// NewDownloadManager creates a new DownloadManager
func NewDownloadManager(downloadDir string) (*DownloadManager, error) {
	// Initialize download directory, create if not exists
	// (Error handling omitted for brevity)
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
	fmt.Printf("Download added with ID %d\n", id)
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

func (mgr *DownloadManager) Watch() {
	reader := bufio.NewReader(os.Stdin)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				mgr.ListDownloads()
			case <-done:
				return
			}
		}
	}()

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		if strings.TrimSpace(input) == "q" {
			break
		}
	}

	done <- true
	fmt.Println("Exited watch mode.")
}
