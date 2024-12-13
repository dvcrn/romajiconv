.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	go vet ./...
	go fmt ./...

.PHONY: all
all: lint test
