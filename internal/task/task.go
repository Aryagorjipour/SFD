package task

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// DownloadTask represents a download task
type DownloadTask struct {
	ID          int
	URL         string
	Status      string
	Progress    float64
	downloadDir string

	ctx        context.Context
	cancelFunc context.CancelFunc

	mu         sync.Mutex
	paused     bool
	resumeChan chan struct{}
}

// NewDownloadTask creates a new DownloadTask
func NewDownloadTask(id int, url, downloadDir string) *DownloadTask {
	ctx, cancel := context.WithCancel(context.Background())
	return &DownloadTask{
		ID:          id,
		URL:         url,
		Status:      "Queued",
		Progress:    0.0,
		downloadDir: downloadDir,
		ctx:         ctx,
		cancelFunc:  cancel,
		resumeChan:  make(chan struct{}),
	}
}

// Start starts the download task
func (dt *DownloadTask) Start() {
	dt.mu.Lock()
	dt.Status = "Downloading"
	dt.mu.Unlock()

	// Extract file name from URL
	parts := strings.Split(dt.URL, "/")
	fileName := parts[len(parts)-1]
	filePath := filepath.Join(dt.downloadDir, fileName)

	// Open file for writing (supports resuming)
	var file *os.File
	var err error
	var downloaded int64 = 0

	if _, err = os.Stat(filePath); err == nil {
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			dt.updateStatus("Error")
			fmt.Printf("Error opening file %s: %v\n", filePath, err)
			return
		}
		fi, err := file.Stat()
		if err == nil {
			downloaded = fi.Size()
		}
	} else {
		file, err = os.Create(filePath)
		if err != nil {
			dt.updateStatus("Error")
			fmt.Printf("Error creating file %s: %v\n", filePath, err)
			return
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file %s: %v\n", filePath, err)
		}
	}(file)

	// Create HTTP request with Range header for resuming
	req, err := http.NewRequestWithContext(dt.ctx, "GET", dt.URL, nil)
	if downloaded > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", downloaded))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		dt.updateStatus("Error")
		fmt.Printf("Error downloading %s: %v\n", dt.URL, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
		dt.updateStatus("Error")
		fmt.Printf("Server returned status %s for %s\n", resp.Status, dt.URL)
		return
	}

	// Determine total size
	var total int64
	if resp.Header.Get("Content-Length") != "" {
		cl, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err == nil {
			total = cl + downloaded
		}
	}

	buf := make([]byte, 32*1024) // 32KB buffer
	for {
		select {
		case <-dt.ctx.Done():
			dt.updateStatus("Canceled")
			return
		default:
			// Handle pausing
			dt.mu.Lock()
			if dt.paused {
				dt.mu.Unlock()
				dt.updateStatus("Paused")
				<-dt.resumeChan
				dt.mu.Lock()
				dt.paused = false
				dt.Status = "Downloading"
				dt.mu.Unlock()
			} else {
				dt.mu.Unlock()
			}

			n, err := resp.Body.Read(buf)
			if n > 0 {
				_, writeErr := file.Write(buf[:n])
				if writeErr != nil {
					dt.updateStatus("Error")
					fmt.Printf("Error writing to file %s: %v\n", filePath, writeErr)
					return
				}
				downloaded += int64(n)
				if total > 0 {
					dt.updateProgress((float64(downloaded) / float64(total)) * 100)
				}
			}
			if err != nil {
				if err == io.EOF {
					dt.updateProgress(100.0)
					dt.updateStatus("Completed")
				} else {
					dt.updateStatus("Error")
					fmt.Printf("Error downloading %s: %v\n", dt.URL, err)
				}
				return
			}
		}
	}
}

// Pause pauses the download task
func (dt *DownloadTask) Pause() {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if dt.Status != "Downloading" {
		fmt.Printf("Download ID %d is not in progress.\n", dt.ID)
		return
	}
	dt.paused = true
}

// Resume resumes the download task
func (dt *DownloadTask) Resume() {
	dt.mu.Lock()
	defer dt.mu.Unlock()
	if dt.Status != "Paused" {
		fmt.Printf("Download ID %d is not paused.\n", dt.ID)
		return
	}
	dt.resumeChan <- struct{}{}
}

// Cancel cancels the download task
func (dt *DownloadTask) Cancel() {
	dt.cancelFunc()
}

func (dt *DownloadTask) updateProgress(p float64) {
	dt.mu.Lock()
	dt.Progress = p
	dt.mu.Unlock()
}

func (dt *DownloadTask) updateStatus(status string) {
	dt.mu.Lock()
	dt.Status = status
	dt.mu.Unlock()
}
