package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
	"github.com/Aryagorjipour/smart-file-downloader/internal/tui"
	"github.com/Aryagorjipour/smart-file-downloader/internal/ui"
)

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
	// Define flags
	downloadDir := flag.String("dir", "./downloads", "Download directory path")
	legacyUI := flag.Bool("legacy-ui", false, "Use legacy command-line UI instead of TUI")
	flag.Parse()

	// Ensure download directory exists
	err := ensureDownloadDir(*downloadDir)
	if err != nil {
		fmt.Printf("error creating download directory: %v\n", err)
		os.Exit(1)
	}

	// Initialize download manager
	mgr, err := manager.NewDownloadManager(*downloadDir)
	if err != nil {
		fmt.Printf("Error initializing download manager: %v\n", err)
		os.Exit(1)
	}

	// Start appropriate UI
	if *legacyUI {
		ui.StartLegacy(mgr)
	} else {
		// Try to start TUI, fallback to legacy on error
		if err := tui.Start(mgr); err != nil {
			fmt.Printf("TUI error: %v\n", err)
			fmt.Println("Falling back to legacy UI...")
			ui.StartLegacy(mgr)
		}
	}
}
