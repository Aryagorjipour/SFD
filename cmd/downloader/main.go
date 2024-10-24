package main

import (
	"fmt"
	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	"github.com/Aryagorjipour/smart-file-downloader/internal/ui"
	"os"
)

func ensureDownloadDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
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
		fmt.Printf("error creating download directory: %v", err)
		os.Exit(1)
	}

	mgr, err := manager.NewDownloadManager(downloadDir)
	if err != nil {
		fmt.Printf("Error initializing download manager: %v\n", err)
		os.Exit(1)
	}

	ui.Start(mgr)
}
