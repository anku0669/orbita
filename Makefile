.PHONY: test build run fmt vet

BINARY ?= orbita

fmt:
	gofmt -w ./cmd ./internal

test:
	go test ./...

vet:
	go vet ./...

build:
	mkdir -p dist
	go build -trimpath -ldflags='-s -w' -o dist/$(BINARY) ./cmd/orbita

run:
	go run ./cmd/orbita --help
