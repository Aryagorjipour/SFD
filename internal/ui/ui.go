package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Aryagorjipour/smart-file-downloader/internal/manager"
)

// Start starts the UI
func Start(mgr *manager.DownloadManager) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Concurrent File Downloader")
	fmt.Printf("Downloads will be saved to: %s\n", mgr.DownloadDir)

	for {
		fmt.Println("Commands:")
		fmt.Println("1. add <URL>")
		fmt.Println("2. list")
		fmt.Println("3. pause <ID>")
		fmt.Println("4. resume <ID>")
		fmt.Println("5. cancel <ID>")
		fmt.Println("6. watch")
		fmt.Println("0. exit")
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		parts := strings.Split(input, " ")
		command := strings.ToLower(parts[0])

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <URL>")
				continue
			}
			url := parts[1]
			mgr.AddDownload(url)
		case "list":
			mgr.ListDownloads()
		case "pause":
			if len(parts) < 2 {
				fmt.Println("Usage: pause <ID>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			mgr.PauseDownload(id)
		case "resume":
			if len(parts) < 2 {
				fmt.Println("Usage: resume <ID>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			mgr.ResumeDownload(id)
		case "cancel":
			if len(parts) < 2 {
				fmt.Println("Usage: cancel <ID>")
				continue
			}
			id, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			mgr.CancelDownload(id)
		case "watch":
			mgr.Watch()
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}
}
