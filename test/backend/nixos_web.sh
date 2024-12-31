env -i \
  GOARCH="$GOARCH" \
  GOPATH="$GOPATH" \
  GOROOT="$GOROOT" \
  GOCACHE="$GOCACHE" \
  USER="$USER" \
  XDG_CACHE_HOME="$XDG_CACHE_HOME" \
  HOME="$HOME" \
  PATH="$GOPATH/bin:$PATH" \
  "$GOPATH/bin/wasmserve" ./test/backend
