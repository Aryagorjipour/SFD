package main

import (
	"encoding/json"
	"fmt"
	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	"net/http"
	"os"
)

type DownloadRequest struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func ensureDownloadDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		mode := int(0750)
		err := os.MkdirAll(dir, os.FileMode(mode))
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %v", dir, err)
		}
	}
	return nil
}

func main() {
	downloadDir := "./downloads"

	if len(os.Args) > 1 {
		downloadDir = os.Args[1]
	}

	err := ensureDownloadDir(downloadDir)
	if err != nil {
		fmt.Printf("Error creating download directory: %v", err)
		os.Exit(1)
	}

	mgr, err := manager.NewDownloadManager(downloadDir)
	if err != nil {
		fmt.Printf("Error initializing download manager: %v\n", err)
		os.Exit(1)
	}

	// Endpoint برای شروع دانلود
	http.HandleFunc("/start-download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req DownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		id := mgr.AddDownload(req.URL)
		if id <= 0 {
			http.Error(w, "Failed to start download: "+req.URL, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf("Download started successfully with ID %d", id)))
		if err != nil {
			http.Error(w, "Failed to start download: "+req.URL, http.StatusInternalServerError)
			return
		}
	})

	// Endpoint برای دریافت وضعیت دانلود
	http.HandleFunc("/download-status", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		statuses := mgr.GetAllDownloadStatuses()
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(statuses)
		if err != nil {
			http.Error(w, "Could not encode request", http.StatusMethodNotAllowed)
			return
		}
	})

	// Endpoint برای مکث دانلود
	http.HandleFunc("/pause-download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req DownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		mgr.PauseDownload(req.ID)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf("Download %d paused successfully", req.ID)))
		if err != nil {
			http.Error(w, "Failed to pause download: "+req.URL, http.StatusInternalServerError)
			return
		}
	})

	// Endpoint برای از سرگیری دانلود
	http.HandleFunc("/resume-download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req DownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		mgr.ResumeDownload(req.ID)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf("Download %d resumed successfully", req.ID)))
		if err != nil {
			http.Error(w, "Failed to resume download: "+req.URL, http.StatusInternalServerError)
			return
		}
	})

	// Endpoint برای لغو دانلود
	http.HandleFunc("/cancel-download", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req DownloadRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		mgr.CancelDownload(req.ID)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf("Download %d cancelled successfully", req.ID)))
		if err != nil {
			http.Error(w, "Failed to cancel download: "+req.URL, http.StatusInternalServerError)
			return
		}
	})

	// Endpoint برای حذف دانلودهای لغو شده
	http.HandleFunc("/clear-cancelled", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		mgr.ClearCancelledDownloads()
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("Cancelled downloads cleared successfully"))
		if err != nil {
			http.Error(w, "Failed to clear cancelled: ", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
