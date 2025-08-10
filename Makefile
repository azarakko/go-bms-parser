# Makefile for go-bms-parser

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD) fmt
GOLINT=golangci-lint

# All source files
GO_FILES := $(shell find . -name '*.go' -not -path "./vendor/*")

# Default target
.PHONY: all
all: help

# Help
.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  deps          Install dependencies"
	@echo "  fmt           Format all go source files"
	@echo "  lint          Run linter"
	@echo "  test          Run tests with coverage"
	@echo "  build         Build the application"
	@echo "  run           Run the CLI application"
	@echo "  ci            Run all CI checks (fmt, lint, test, build)"

# Dependencies
.PHONY: deps
deps:
	$(GOCMD) mod download
	$(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Format
.PHONY: fmt
fmt:
	@echo "Formatting go files..."
	@$(GOFMT) -w $(GO_FILES)

# Lint
.PHONY: lint
lint:
	@echo "Running linter..."
	@$(GOLINT) run ./...

# Test
.PHONY: test
test:
	@echo "Running tests..."
	@$(GOTEST) -v -race -coverprofile=coverage.out ./...

# Build
.PHONY: build
build:
	@echo "Building application..."
	@$(GOBUILD) -v -o build/bms-parser-cli ./cmd/bms-parser-cli

# Run
.PHONY: run
run: build
	@echo "Running application..."
	@./build/bms-parser-cli

# Lint-Fmt
.PHONY: lint-fmt
lint-fmt:
	@echo "Checking go files format..."
	@test -z $(shell $(GOFMT) -l $(GO_FILES)) || (echo "Go files are not formatted. Please run 'make fmt'"; exit 1)

# CI
.PHONY: ci
ci: lint-fmt lint test build

.PHONY: clean
clean:
	@echo "Cleaning up..."
	@rm -f coverage.out
	@rm -rf build/
