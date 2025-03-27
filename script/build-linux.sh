#!/bin/bash

# Set paths
SRC="main.go"
OUT_DIR="dist"

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64) ARCH="arm64" ;;
    armv7l) ARCH="arm" ;;
    armhf) ARCH="arm" ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Ensure the output directory exists
mkdir -p "$OUT_DIR"

# Set output file name
OUT="$OUT_DIR/branded-linux-$ARCH"

# Build for detected architecture
echo "Detected architecture: $ARCH"
echo "Building for Linux-$ARCH..."

GOOS=linux GOARCH=$ARCH go build -o "$OUT" "$SRC"

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "Build successful: $OUT"
else
    echo "Build failed."
    exit 1
fi
