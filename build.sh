#!/bin/bash
set -e

echo "Building ccat v2.0"
mkdir -p build

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -trimpath -mod=vendor -x -o build/ccat-linux-amd64 ./main.go
GOOS=linux GOARCH=arm GOARM=5 go build -trimpath -mod=vendor -x -o build/ccat-linux-armv5 ./main.go
GOOS=linux GOARCH=arm64 go build -trimpath -mod=vendor -x -o build/ccat-linux-arm64 ./main.go
GOOS=android GOARCH=arm64 go build -trimpath -mod=vendor -x -o build/ccat-android-arm64 ./main.go
GOOS=windows GOARCH=amd64 go build -trimpath -mod=vendor -x -o build/ccat-windows-amd64.exe ./main.go
GOOS=darwin GOARCH=amd64 go build -trimpath -mod=vendor -x -o build/ccat-darwin-amd64 ./main.go

echo "Build complete. Binaries in build/ directory."
