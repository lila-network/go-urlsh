# Set the default Go build flags
GOFLAGS = -ldflags='-w -s -X github.com/lila-network/go-urlsh/misc/constants.Version=$(VERSION)'

PROJECT=go-urlsh

.PHONY: test

# Compile the application into the bin/ folder
build:
	go build $(GOFLAGS) -o bin/$(PROJECT) cmd/$(PROJECT)/main.go

# Run the application
run:
	go run cmd/$(PROJECT)/main.go

# Clean the build artifacts
clean:
	rm -rf bin

# Run go tests
test:
	go test ./...

# Set a version for the build
VERSION := $(shell git describe --tags --always)
