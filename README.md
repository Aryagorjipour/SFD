# SFD (Smart File Downloader)

![License](https://img.shields.io/github/license/Aryagorjipour/SFD)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/Aryagorjipour/SFD/ci.yml)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/Aryagorjipour/SFD)

A command-line application written in Go that allows users to download multiple files concurrently with a modern Terminal UI (TUI), featuring batch URL input, automatic cleanup, and real-time progress tracking.

## Features

- **Modern TUI:** Beautiful terminal interface with mouse support and keyboard shortcuts
- **Batch URL Input:** Paste multiple URLs at once (newline-separated)
- **Smart URL Display:** Automatically trims query parameters for cleaner viewing
- **Auto-Cleanup:** Completed, canceled, and errored downloads are automatically removed
- **Concurrent Downloads:** Download multiple files simultaneously using goroutines
- **Real-time Progress:** Live progress bars and status updates
- **Pause/Resume/Cancel:** Full control over individual downloads
- **Configurable:** Customize download directory via flags

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/Aryagorjipour/SFD.git
   cd SFD
   ```

2. **Build the Application:**
   ```bash
   ./scripts/build.sh
   ```
   The executable will be located at ./bin/downloader.

## Usage

### Modern TUI (Default)

Run the application to launch the Terminal UI:

```bash
./bin/downloader
```

With custom download directory:
```bash
./bin/downloader --dir ./my_downloads
```

#### TUI Keyboard Shortcuts

- **`a`** - Add batch URLs (paste multiple URLs, one per line)
- **`↑/↓` or `k/j`** - Navigate between downloads
- **`p`** - Pause selected download
- **`r`** - Resume selected download
- **`c`** - Cancel selected download
- **`q` or `Ctrl+C`** - Quit application

#### Adding Batch URLs

1. Press `a` to enter batch input mode
2. Paste your URLs (one per line), for example:
   ```
   https://example.com/file1.mp4
   https://example.com/file2.mp4
   https://example.com/file3.mp4
   ```
3. Press `Ctrl+S` to submit
4. Press `Esc` to cancel

#### Features

- **Automatic URL Trimming:** Long URLs with query parameters are automatically cleaned
  - Before: `https://example.com/file.mp4?md5=xxx&u=yyy&expires=zzz`
  - After: `https://example.com/file.mp4`
- **Auto-Cleanup:** Completed, canceled, and errored downloads disappear automatically
- **Real-time Updates:** Download progress refreshes every second
- **Mouse Support:** Click buttons (coming soon)

### Legacy CLI Mode

For the classic command-line interface:

```bash
./bin/downloader --legacy-ui
```

#### Legacy Commands

- `add <URL>` - Download a file
- `list` - List all downloads
- `pause <ID>` - Pause a download
- `resume <ID>` - Resume a download
- `cancel <ID>` - Cancel a download
- `watch` - Auto-refresh download list
- `exit` - Exit the application

## Configuration

Modify the `configs/config.yaml` file to customize settings:
```yaml
download_directory: "./downloads"
max_concurrent_downloads: 5
```

## Testing

Run tests using the following command:
```bash
go test ./tests/...
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

MIT License © 2024 Aryagorjipour
