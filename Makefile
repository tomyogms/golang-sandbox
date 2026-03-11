.PHONY: help docker-prod docker-dev docker-stop docker-logs docker-test dev-local run-local build test test-v test-cov fmt lint clean

help:
	@echo "Docker Commands (Recommended):"
	@echo "  make docker-prod       - Run production server in Docker"
	@echo "  make docker-dev        - Run development server in Docker with live reload"
	@echo "  make docker-stop       - Stop all Docker containers"
	@echo "  make docker-logs       - View Docker container logs"
	@echo "  make docker-test       - Run tests in Docker container"
	@echo ""
	@echo "Local Commands (Alternative - requires Go 1.25+):"
	@echo "  make dev-local         - Run with live reload locally (requires air)"
	@echo "  make run-local         - Run without live reload locally"
	@echo "  make build             - Build binary locally"
	@echo "  make test              - Run tests locally"
	@echo "  make test-v            - Run tests locally with verbose output"
	@echo "  make test-cov          - Generate coverage report locally"
	@echo ""
	@echo "Code Quality:"
	@echo "  make fmt               - Format code with gofmt"
	@echo "  make lint              - Run go vet"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean             - Remove build artifacts and containers"

# ============================================================================
# DOCKER COMMANDS (Recommended)
# ============================================================================

docker-prod:
	@echo "🐳 Starting production server in Docker..."
	docker-compose up --build

docker-dev:
	@echo "🐳 Starting development server in Docker with live reload..."
	docker-compose -f docker-compose.dev.yml up --build

docker-stop:
	@echo "🛑 Stopping Docker containers..."
	docker-compose down

docker-logs:
	@echo "📋 Viewing Docker logs..."
	docker-compose logs -f api

docker-test:
	@echo "🧪 Running tests in Docker..."
	docker-compose exec api go test ./...

# ============================================================================
# LOCAL COMMANDS (Alternative)
# ============================================================================

dev-local:
	@echo "🚀 Starting local development server with live reload..."
	@command -v air >/dev/null 2>&1 && air || ~/go/bin/air

run-local:
	@echo "🚀 Starting local server..."
	go run cmd/api/main.go

build:
	@echo "🔨 Building binary..."
	go build -o api cmd/api/main.go
	@echo "✅ Binary created: ./api"

# ============================================================================
# TEST COMMANDS
# ============================================================================

test:
	@echo "🧪 Running tests..."
	go test ./...

test-v:
	@echo "🧪 Running tests (verbose)..."
	go test ./... -v

test-cov:
	@echo "📊 Running tests with coverage..."
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report: coverage.html"

# ============================================================================
# CODE QUALITY COMMANDS
# ============================================================================

fmt:
	@echo "🎨 Formatting Go code..."
	go fmt ./...
	@echo "✅ Code formatted"

lint:
	@echo "🔍 Running go vet..."
	go vet ./...
	@echo "✅ Lint check complete"

# ============================================================================
# CLEANUP
# ============================================================================

clean:
	@echo "🧹 Cleaning up..."
	rm -f api coverage.out coverage.html
	rm -rf tmp/
	docker-compose down 2>/dev/null || true
	docker-compose -f docker-compose.dev.yml down 2>/dev/null || true
	@echo "✅ Clean complete"
