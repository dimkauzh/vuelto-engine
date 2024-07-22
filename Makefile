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
	GOPROXY=proxy.golang.org go list -m vuelto.me@$(VERSION)

ci_check:
	go build -o bin/test/test test/test/test.go

	go build -o bin/examples/basic-window examples/basic-window/main.go
	go build -o bin/examples/rectangle examples/rectangle/main.go
	go build -o bin/examples/images examples/images/main.go
	go build -o bin/examples/two-windows examples/two-windows/main.go

test:
	go run test/test/test.go

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

