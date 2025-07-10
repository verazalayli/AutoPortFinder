#!/bin/bash

set -e

BINARY_NAME="apf"

# Determine OS and architecture
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
INSTALL_DIR="/usr/local/bin"

# Set file extension for Windows
EXT=""
if [ "$GOOS" == "windows" ]; then
  EXT=".exe"
  INSTALL_DIR="$HOME/bin"  # Windows: use user-local bin directory
  mkdir -p "$INSTALL_DIR"
fi

FULL_BINARY_NAME="${BINARY_NAME}${EXT}"

echo "➡️ Detected platform: $GOOS/$GOARCH"
echo "➡️ Building the project for $GOOS/$GOARCH..."
GOOS=$GOOS GOARCH=$GOARCH go build -o "$FULL_BINARY_NAME"

echo "➡️ Moving binary to $INSTALL_DIR (sudo may be required on Unix systems)..."
if [[ "$GOOS" == "windows" ]]; then
  mv "$FULL_BINARY_NAME" "$INSTALL_DIR/$FULL_BINARY_NAME"
else
  sudo mv "$FULL_BINARY_NAME" "$INSTALL_DIR/$FULL_BINARY_NAME"
  sudo chmod +x "$INSTALL_DIR/$FULL_BINARY_NAME"
fi

# Optionally add to PATH for current session (Windows)
if [[ "$GOOS" == "windows" ]]; then
  export PATH="$INSTALL_DIR:$PATH"
  echo "📎 Make sure $INSTALL_DIR is in your system PATH."
fi

echo "✅ Installation complete. Verifying..."
echo
"$FULL_BINARY_NAME" --help || echo "⚠️ Command installed, but '--help' is not supported."

echo
echo "📌 You can now use the command: $FULL_BINARY_NAME"
echo "Example: $FULL_BINARY_NAME --format=table --process=nginx"
