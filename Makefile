.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: format
fmt:
	go fmt ./...

.PHONY: test
test:
	go test -v ./...
