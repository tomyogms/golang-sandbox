# Golang Sandbox

A clean, production-ready Go boilerplate with best practices, proper separation of concerns, and comprehensive tests.

## Features

- ✨ **Single Health Check Endpoint** - `/up` endpoint for service availability checks
- 🏗️ **Clean Architecture** - Separation of concerns with config, server, and HTTP packages
- 📋 **Structured Logging** - JSON-formatted logging using Go's `log/slog`
- 🧪 **Unit Tests** - Comprehensive test coverage for all packages
- ⚡ **Graceful Shutdown** - Proper HTTP server lifecycle management
- 🔧 **Environment Configuration** - 12-factor app configuration from environment variables

## Prerequisites

- Go 1.25.0 or later
- `curl` (for testing the API)
- Docker & Docker Compose (optional, for containerized deployment)

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── config/
│   │   ├── config.go               # Configuration management
│   │   └── config_test.go          # Configuration tests
│   ├── server/
│   │   ├── server.go               # HTTP server lifecycle
│   │   └── server_test.go          # Server tests
│   └── http/
│       ├── router.go               # HTTP handlers & routing
│       └── router_test.go          # Handler tests
├── .air.toml                       # Air live reload configuration
├── .dockerignore                   # Docker build ignore patterns
├── .gitignore                      # Git ignore rules
├── dev.sh                          # Development server script
├── docker-compose.yml              # Production docker-compose
├── docker-compose.dev.yml          # Development docker-compose with live reload
├── Dockerfile                      # Multi-stage Docker build
├── DOCKER.md                       # Docker & Docker Compose guide
├── DOCKER_COMMANDS.md              # Quick Docker commands reference
├── Makefile                        # Development tasks
├── RELEASES.md                     # Changelog and release notes
├── go.mod                          # Go module definition
├── go.sum                          # Go module checksums
└── README.md                       # This file
```

## Installation

1. Clone the repository:
```bash
cd /Users/tomyo/git/golang-sandbox
```

2. Ensure Docker & Docker Compose are installed:
   - [Docker Desktop](https://www.docker.com/products/docker-desktop) includes both

## Running the Application (Docker)

### 🐳 Production Build & Run

For production-like environment with healthcheck:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

View logs:
```bash
docker-compose logs -f api
```

Stop services:
```bash
docker-compose down
```

### 🐳 Development Build with Live Reload

For development with automatic app restart on code changes:

```bash
docker-compose -f docker-compose.dev.yml up --build
```

The app will auto-restart whenever you save `.go` files inside the container.

View logs:
```bash
docker-compose -f docker-compose.dev.yml logs -f api
```

Stop services:
```bash
docker-compose -f docker-compose.dev.yml down
```

### Test the API

```bash
curl http://localhost:8080/up
```

Response:
```json
{"status":"up"}
```

For detailed Docker instructions, see [DOCKER.md](DOCKER.md) and [DOCKER_COMMANDS.md](DOCKER_COMMANDS.md).

## Local Development (Alternative)

If you prefer to run locally without Docker, follow the instructions below.

### Prerequisites for Local Development
- Go 1.25.0 or later
- `air` for live reload: `go install github.com/air-verse/air@latest`

### Running Locally with Live Reload

```bash
# If air is in your PATH
air

# Or using make
make dev-local

# Or using the script
./dev.sh
```

The app will automatically restart whenever you save changes to `.go` files.

Custom port:
```bash
PORT=3000 air
```

### Running Locally without Live Reload

```bash
make run-local
# Or
go run cmd/api/main.go
```

Custom port:
```bash
PORT=3000 go run cmd/api/main.go
```

### Building Local Binary

```bash
make build
```

Output: `./api`

### Local Testing

```bash
make test          # Run all tests
make test-v        # Run tests with verbose output
make test-cov      # Generate coverage report
```

### Local Formatting & Linting

```bash
make fmt           # Format code
make lint          # Run go vet
make clean         # Remove build artifacts
```

## API Endpoints

### Simple Health Check

```bash
curl http://localhost:8080/up
```

Response:
```json
{"status":"up"}
```

### Detailed Health Check (with Database Status)

```bash
curl http://localhost:8080/health
```

Response (when database is healthy):
```json
{"status":"healthy","database":"healthy"}
```

Response (when database is unhealthy):
```json
{"status":"unhealthy","database":"unhealthy"}
```

Status codes:
- `200 OK` - All services healthy
- `503 Service Unavailable` - Database or other services unhealthy

## Running Tests with Docker

```bash
# Run tests in production container
docker-compose exec api go test ./...

# Run tests with verbose output
docker-compose exec api go test ./... -v

# Run tests with coverage
docker-compose exec api go test ./... -coverprofile=coverage.out
```

## Running Tests Locally

### All Tests
```bash
go test ./...
```

### Verbose Output
```bash
go test ./... -v
```

### With Coverage Report
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Make Commands

The project includes a [Makefile](Makefile) with convenient commands:

```bash
make help              # Show all available commands

# Docker commands (recommended)
make docker-prod       # Run production container
make docker-dev        # Run development container with live reload
make docker-stop       # Stop all containers
make docker-logs       # View container logs
make docker-test       # Run tests in container

# Local commands (alternative)
make dev-local         # Run with live reload locally
make run-local         # Run locally without reload
make build             # Build binary
make test              # Run tests locally
make test-v            # Run tests with verbose output
make test-cov          # Generate coverage report
make fmt               # Format code
make lint              # Run linter
make clean             # Clean build artifacts
```

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port number |
| `ENV` | `development` | Environment name (development, production, etc.) |

## Configuration

Configuration is loaded from environment variables with sensible defaults. The application uses the 12-factor app methodology.

See [internal/config/config.go](internal/config/config.go) for implementation details.

## Development

### Code Style

The project follows standard Go conventions:
- `gofmt` for code formatting
- `go vet` for potential bugs
- Exported functions and types have comments

### Adding New Endpoints

1. Add a handler method to the `Server` struct in [internal/http/router.go](internal/http/router.go)
2. Register the route in the `NewServer` function
3. Add tests in [internal/http/router_test.go](internal/http/router_test.go)

Example:
```go
func (s *Server) handleExample(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"message": "example"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}
```

Then register in `NewServer`:
```go
mux.HandleFunc("GET /example", s.handleExample)
```

## Troubleshooting

### Port Already in Use

If port 8080 is already in use:
```bash
PORT=3000 go run cmd/api/main.go
```

### Module Not Found

Clear the module cache and re-download:
```bash
go clean -modcache
go mod download
```

## License

MIT
