#!/bin/bash

# Build the downloader executable
go build -o bin/downloader ./cmd/downloader

echo "Build completed. Executable is located at ./bin/downloader"
