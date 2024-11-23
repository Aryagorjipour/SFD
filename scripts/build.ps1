# Set the output directories
$OUT_DIR = "./dist"
$ELECTRON_BIN_DIR = "./electron-gui/binaries"
$ELECTRON_OUT_DIR = "./electron-gui/release-builds"

# Create output directories if they don't exist
If (-Not (Test-Path -Path $OUT_DIR)) {
    New-Item -ItemType Directory -Path $OUT_DIR
}
If (-Not (Test-Path -Path $ELECTRON_BIN_DIR)) {
    New-Item -ItemType Directory -Path $ELECTRON_BIN_DIR
}
If (-Not (Test-Path -Path $ELECTRON_OUT_DIR)) {
    New-Item -ItemType Directory -Path $ELECTRON_OUT_DIR
}

# Build the Go application for Windows
Write-Output "Building Go application for Windows..."
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "$OUT_DIR/sfd-windows.exe" ./cmd/downloader
If ($LASTEXITCODE -ne 0) {
    Write-Output "Failed to build for Windows"
    exit 1
}

# Build the Go application for Linux
Write-Output "Building Go application for Linux..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "$OUT_DIR/sfd-linux" ./cmd/downloader
If ($LASTEXITCODE -ne 0) {
    Write-Output "Failed to build for Linux"
    exit 1
}

Write-Output "Go application build completed successfully."

# Copy Go binaries to Electron project
Write-Output "Copying Go binaries to Electron project..."
Copy-Item "$OUT_DIR/sfd-windows.exe" -Destination $ELECTRON_BIN_DIR
Copy-Item "$OUT_DIR/sfd-linux" -Destination $ELECTRON_BIN_DIR

# Navigate to the Electron project directory
Set-Location -Path "electron-gui"

# Install dependencies if not installed
If (-Not (Test-Path -Path "node_modules")) {
    Write-Output "Installing Node.js dependencies..."
    npm install
}

# Build the Electron application to create an installer for Windows and Linux
Write-Output "Building Electron application for Windows and Linux..."
npx electron-builder -wl
If ($LASTEXITCODE -ne 0) {
    Write-Output "Electron build failed"
    exit 1
}

Write-Output "Electron build completed. Output is located at $ELECTRON_OUT_DIR"

# Navigate back to root directory
Set-Location -Path ".."

Write-Output "Build process completed successfully."
