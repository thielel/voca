.PHONY: build run test clean tidy seed

# Build the application
build:
	go build -o bin/api ./cmd/api

# Run the application
run:
	go run ./cmd/api

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Tidy dependencies
tidy:
	go mod tidy

# Download dependencies
deps:
	go mod download

# Seed database with example data
seed:
	go run ./cmd/seed

# Format code
fmt:
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	golangci-lint run

# Build for production
build-prod:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/api ./cmd/api

# Run with live reload (requires air)
dev:
	air

# Help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage"
	@echo "  clean        - Clean build artifacts"
	@echo "  tidy         - Tidy dependencies"
	@echo "  deps         - Download dependencies"
	@echo "  seed         - Seed database with example data"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  build-prod   - Build for production"
	@echo "  dev          - Run with live reload"
