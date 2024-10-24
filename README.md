# Concurrent File Downloader

A command-line application written in Go that allows users to download multiple files concurrently with features like pausing, resuming, and canceling downloads.

## Features

- **Concurrent Downloads:** Download multiple files simultaneously using goroutines.
- **Progress Tracking:** Real-time display of download progress for each file.
- **Pause/Resume:** Ability to pause and resume individual downloads.
- **Cancel Downloads:** Option to cancel ongoing downloads.
- **Configurable:** Customize download directory and maximum concurrent downloads via configuration.

## Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/concurrent-file-downloader.git
   cd concurrent-file-downloader
   ```

2. **Build the Application:**
   ```bash
   ./scripts/build.sh
   ```
   The executable will be located at ./bin/downloader.

## Usage
   ```bash
   ./bin/downloader [download_directory]
   ```
- _Example_:
   ```bash
   ./bin/downloader ./my_downloads
   ```
- ### Commands
    - Download a file:
    ```bash
    download <URL>
    ```
    - List all downloads:
    ```bash
    list
    ```
    - Pause a download:
    ```bash
    pause <ID>
    ```
    - Resume a download:
    ```bash
    resume <ID>
    ```
    - Cancel a download:
    ```bash
    cancel <ID>
    ```
   - Exit the application:
   ```bash
    exit
   ```

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