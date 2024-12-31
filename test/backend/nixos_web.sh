GO_BINARY=$(which go)
if [ -z "$GO_BINARY" ]; then
  echo "Error: Go binary not found. Please ensure Go is installed and available in your PATH."
  exit 1
fi

GOROOT=$(dirname $(dirname "$GO_BINARY"))
if [ ! -d "$GOROOT" ]; then
  echo "Error: GOROOT directory $GOROOT does not exist."
  exit 1
fi

CACHE_DIR="/home/$USER/.cache/go-build"

env -i \
  GOARCH="amd64" \
  GOPATH="$HOME/.local/go" \
  GOROOT="$GOROOT/share/go" \
  GOCACHE="$CACHE_DIR" \
  USER="$USER" \
  XDG_CACHE_HOME="$HOME/.cache" \
  HOME="/home/$USER" \
  PATH="$HOME/.local/go/bin:$PATH" \
  "$GOPATH/bin/wasmserve" ./test/backend
