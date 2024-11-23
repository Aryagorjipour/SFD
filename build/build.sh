# Set the output directories
OUT_DIR="./build/release-builds"
BINARY_DIR="./frontend/binaries"
ELECTRON_OUT_DIR="./build/release-builds/electron"

# Create output directories if they don't exist
mkdir -p $OUT_DIR
mkdir -p $BINARY_DIR
mkdir -p $ELECTRON_OUT_DIR

# Navigate to backend directory to build Go application
cd backend

# Build the Go application for Windows and Linux
echo "Building Go application for Windows..."
GOOS=windows GOARCH=amd64 go build -o ../$OUT_DIR/sfd-windows.exe ./cmd/downloader
if [ $? -ne 0 ]; then
    echo "Failed to build for Windows"
    exit 1
fi

echo "Building Go application for Linux..."
GOOS=linux GOARCH=amd64 go build -o ../$OUT_DIR/sfd-linux ./cmd/downloader
if [ $? -ne 0 ]; then
    echo "Failed to build for Linux"
    exit 1
fi

echo "Go application build completed successfully."

# Navigate back to root directory
cd ..

# Copy Go binaries to the binaries folder in the frontend project
echo "Copying Go binaries to frontend binaries folder..."
cp $OUT_DIR/sfd-windows.exe $BINARY_DIR/
cp $OUT_DIR/sfd-linux $BINARY_DIR/

# Continue with Electron build process
cd frontend/electron-gui

# Install dependencies if not installed
if [ ! -d "node_modules" ]; then
    echo "Installing Node.js dependencies..."
    npm install
fi

# Build the Electron application to create an installer for Windows and Linux
echo "Building Electron application for Windows and Linux..."
npx electron-builder -wl
if [ $? -ne 0 ]; then
    echo "Electron build failed"
    exit 1
fi

echo "Electron build completed. Output is located at $ELECTRON_OUT_DIR"

# Navigate back to root directory
cd ../..

echo "Build process completed successfully."
