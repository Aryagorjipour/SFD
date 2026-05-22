# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

SFD (Smart File Downloader) is a concurrent file download manager written in Go. It provides CLI functionality to download multiple files simultaneously with pause/resume/cancel capabilities.

## Build & Run Commands

```bash
# Build the application
./scripts/build.sh
# Output: ./bin/downloader

# Run directly with go
go run ./cmd/downloader [download_directory]

# Run the built executable
./bin/downloader [download_directory]
```

## Testing

```bash
# Run all tests
go test ./tests/...

# Run tests with verbose output
go test -v ./tests/...

# Build all packages (CI check)
go build ./...
```

## Module Information

- **Module path**: `github.com/Aryagorjipour/smart-file-downloader`
- **Go version**: 1.23.2
- **No external dependencies**: Uses only standard library

## Architecture

### Core Components

**Manager Layer** (`internal/manager/manager.go`):
- `DownloadManager` is the central orchestrator
- Maintains a map of download tasks indexed by ID
- Uses `sync.Mutex` for thread-safe task map access
- Assigns sequential IDs starting from 1
- Methods: `AddDownload`, `ListDownloads`, `PauseDownload`, `ResumeDownload`, `CancelDownload`, `Watch`

**Task Layer** (`internal/task/task.go`):
- `DownloadTask` represents individual download operations
- Uses `context.Context` for cancellation support
- Implements pause/resume via channels (`resumeChan`)
- Supports HTTP range requests for resuming partial downloads
- Downloads to `./[downloadDir]/[filename]` with write permissions 0600
- Uses 32KB buffer for efficient I/O

**UI Layer** (`internal/ui/ui.go`):
- Command-line interface using `bufio.Reader`
- Accepts space-separated commands: `add <URL>`, `list`, `pause <ID>`, `resume <ID>`, `cancel <ID>`, `watch`, `exit`

**Main Entry** (`cmd/downloader/main.go`):
- Creates download directory with 0750 permissions if needed
- Initializes DownloadManager and starts UI
- Accepts optional download directory as first argument

### Concurrency Model

- Each download task runs in its own goroutine (spawned in `AddDownload`)
- Manager uses mutex to protect the shared tasks map
- Individual tasks use their own mutex for status/progress updates
- Context-based cancellation for graceful shutdown
- Channel-based pause/resume mechanism

### Download Flow

1. User adds URL → Manager creates task with unique ID → Spawns goroutine
2. Task extracts filename from URL and creates/opens file
3. Checks existing file size and sets HTTP Range header if resuming
4. Downloads in chunks, updating progress percentage
5. Handles pause via channel blocking, resume via channel signal
6. Cancel via context cancellation

### Watch Mode

The `watch` command in the UI triggers `DownloadManager.Watch()` which:
- Prints download list every 1 second using a ticker
- Runs in a goroutine to allow user input
- Exit by pressing 'q'

## Important Details

- Downloads are saved relative to the current working directory: `./[downloadDir]/[filename]`
- File permissions: directories=0750, files=0600
- The config file `configs/config.yaml` exists but is not currently used by the code
- Logger package (`pkg/logger/logger.go`) exists but is not used in current implementation
