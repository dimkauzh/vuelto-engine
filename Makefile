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
