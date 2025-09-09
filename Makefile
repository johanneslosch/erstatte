# Makefile for erstatte
# Executes the same commands as the GitHub Actions

.PHONY: test test-unit test-integration test-all coverage benchmark lint build build-all clean help

# Default target
all: test lint build

# Tests
test: test-unit test-integration
	@echo "✅ All tests successful"

test-unit:
	@echo "🧪 Running unit tests..."
	go test -v ./tests/unit/...

test-integration:
	@echo "🔗 Running integration tests..."
	go test -v ./tests/integration/...

test-all:
	@echo "🧪 All tests with verbose output..."
	go test -v ./tests/...

coverage:
	@echo "📊 Creating coverage report..."
	go test -coverprofile=coverage.out ./tests/...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage Report: coverage.html"

benchmark:
	@echo "⚡ Running benchmarks..."
	go test -bench=. -benchmem ./tests/benchmarks/...

# Code Quality
lint:
	@echo "🔍 Code linting..."
	golangci-lint run
	go vet ./...
	gofmt -s -l .

fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...
	goimports -w .

# Build
build:
	@echo "🏗️  Building binary..."
	go build -o dist/erstatte ./

build-all:
	@echo "🏗️  Building cross-platform binaries..."
	mkdir -p dist
	# Linux
	GOOS=linux GOARCH=amd64 go build -o dist/erstatte-linux-amd64 ./
	GOOS=linux GOARCH=arm64 go build -o dist/erstatte-linux-arm64 ./
	# Windows
	GOOS=windows GOARCH=amd64 go build -o dist/erstatte-windows-amd64.exe ./
	# macOS
	GOOS=darwin GOARCH=amd64 go build -o dist/erstatte-darwin-amd64 ./
	GOOS=darwin GOARCH=arm64 go build -o dist/erstatte-darwin-arm64 ./
	@echo "✅ Binaries created in dist/"

# Dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod download
	go mod verify

deps-update:
	@echo "📦 Updating dependencies..."
	go get -u ./...
	go mod tidy

# Security
security:
	@echo "🔒 Security scan..."
	gosec ./...

# Cleanup
clean:
	@echo "🧹 Cleaning up..."
	rm -rf dist/
	rm -f coverage.out coverage.html
	go clean

# Development
dev: deps fmt test
	@echo "🚀 Development setup completed"

# CI simulation
ci: deps test lint security build-all
	@echo "🎯 CI pipeline simulated"

# Help
help:
	@echo "Available targets:"
	@echo "  test          - Unit and integration tests"
	@echo "  test-unit     - Unit tests only"
	@echo "  test-integration - Integration tests only"
	@echo "  coverage      - Create coverage report"
	@echo "  benchmark     - Run benchmarks"
	@echo "  lint          - Code linting"
	@echo "  fmt           - Format code"
	@echo "  build         - Binary for current system"
	@echo "  build-all     - Cross-platform binaries"
	@echo "  deps          - Install dependencies"
	@echo "  deps-update   - Update dependencies"
	@echo "  security      - Security scan"
	@echo "  clean         - Clean up"
	@echo "  dev           - Development setup"
	@echo "  ci            - Simulate CI pipeline"
	@echo "  help          - Show this help"
