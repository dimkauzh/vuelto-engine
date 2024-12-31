VERSION = none
MESSAGE = Release version $(VERSION)

.PHONY: release proxy-release ci_check

release:
	git add .
	git commit -m "$(MESSAGE)"
	git tag $(VERSION)
	git push origin $(VERSION)
	make proxy-release

proxy-release:
	GOPROXY=proxy.golang.org go list -m vuelto.pp.ua@$(VERSION)

ci_check:
	go build -o bin/test/test ./test/test/
	go build -o bin/test/backend ./test/backend/

	go build -o bin/examples/basic-window ./examples/basic-window/
	go build -o bin/examples/rectangle ./examples/rectangle/
	go build -o bin/examples/images ./examples/images/
	go build -o bin/examples/two-windows ./examples/two-windows/

test:
	go run test/test/test.go

web_nixos:
	@env -i \
		GOARCH="$(GOARCH)" \
		GOPATH="$(GOPATH)" \
		GOROOT="$(GOROOT)" \
		GOCACHE="$(GOCACHE)" \
		USER="$(USER)" \
		XDG_CACHE_HOME="$(XDG_CACHE_HOME)" \
		HOME="$(HOME)" \
		PATH="$(GOPATH)/bin:$(PATH)" \
		"$(GOPATH)/bin/wasmserve" ./test/backend

format:
	go fmt ./pkg/

	go fmt ./test/test/
	go fmt ./test/backend/

	go fmt ./internal/gl/webgl/
	go fmt ./internal/gl/opengl/
	go fmt ./internal/gl/

	go fmt ./internal/windowing/web/
	go fmt ./internal/windowing/x11/
	go fmt ./internal/windowing/cocoa/
	go fmt ./internal/windowing/win32/
	go fmt ./internal/windowing/wayland/
	go fmt ./internal/windowing/

	go fmt ./examples/images/
	go fmt ./examples/rectangle/
	go fmt ./examples/two-windows/
	go fmt ./examples/basic-window/
