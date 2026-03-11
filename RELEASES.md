# Release Notes

All notable changes to the Golang Sandbox project are documented in this file.

## [Unreleased]
### 🎯 Containerized-First Workflow (March 10, 2026)

#### Changed
- **README.md Restructured**
  - Docker/container commands now primary focus
  - Production Docker deployment highlighted upfront
  - Development Docker with live reload emphasized
  - Local development moved to "Alternative" section

- **Makefile Completely Rebuilt**
  - Docker commands as primary targets (`make docker-prod`, `make docker-dev`, etc.)
  - Local commands renamed to secondary targets (`make dev-local`, `make run-local`)
  - New `make docker-prod` - Production Docker deployment
  - New `make docker-dev` - Development Docker with live reload
  - New `make docker-stop` - Stop all containers
  - New `make docker-logs` - View container logs
  - New `make docker-test` - Run tests in Docker

#### Benefits
✅ Consistent development via Docker
✅ Same environment locally and production
✅ Live reload in Docker works seamlessly
✅ Tests run in containers
✅ Better onboarding for new developers

#### Files Modified
- `README.md` - Docker-first documentation structure
- `Makefile` - Docker commands prioritized

---
### � Docker & Docker Compose Setup (March 10, 2026)

#### Added
- **Docker Multi-Stage Build**
  - `Dockerfile` - Optimized production and build stages
  - Stage 1: Builder - Compiles Go binary on golang:1.25-alpine
  - Stage 2: Runtime - Runs on minimal alpine:3.20 image
  - Final image size: ~15-20 MB
  - Security: Runs as non-root `appuser`

- **Docker Compose Configurations**
  - `docker-compose.yml` - Production deployment
    - Builds and deploys the API container
    - Port 8080 exposed
    - Healthcheck endpoint monitoring (`/up`)
    - Auto-restart policy
  
  - `docker-compose.dev.yml` - Development with live reload
    - Uses builder stage with volume mounts
    - Air installed automatically for live reload
    - Code changes automatically trigger restart
    - Development environment configuration

- **Docker Optimization Files**
  - `.dockerignore` - Excludes unnecessary files from build context
    - Reduces build context size
    - Improves build performance
    - Ignores git, IDE, test, and doc files

- **Comprehensive Docker Documentation**
  - `DOCKER.md` - Complete Docker & Docker Compose guide
    - Quick start commands
    - Production vs development setup
    - Manual Docker image building
    - Multi-stage build explanation
    - Environment variables in Docker
    - Debugging and troubleshooting
    - Performance best practices
    - Registry pushing instructions

  - `DOCKER_COMMANDS.md` - Quick reference for common Docker commands
    - Docker-compose up/down/logs commands
    - API testing from host and container
    - Manual docker build/run commands
    - Container and image management
    - Build optimization tips
    - Common issues and solutions
    - Environment setup
    - Cleanup commands

#### Modified Files
- `Dockerfile` - Updated binary build path and name
  - Changed from `./cmd/api` to `cmd/api/main.go`
  - Updated binary reference to match

- `docker-compose.yml` - Simplified and cleaned up
  - Removed PostgreSQL database service (no longer needed)
  - Removed database environment variables
  - Added healthcheck for container monitoring
  - Kept only API service with clean configuration

- `README.md` - Added Docker information
  - Added Docker & Docker Compose to prerequisites
  - Updated project structure to include Docker files
  - Added Docker quick commands section
  - Link to comprehensive DOCKER.md guide

#### Best Practices Implemented
✅ Multi-stage Docker builds for optimal image size
✅ Alpine base images for security and minimal footprint
✅ Non-root user for container security
✅ Healthcheck monitoring
✅ Proper .dockerignore configuration
✅ Development and production compose files
✅ Clear Docker documentation

#### Files Modified/Created
```
Modified:
- Dockerfile
- docker-compose.yml
- README.md
- RELEASES.md

Created:
- docker-compose.dev.yml
- .dockerignore
- DOCKER.md
- DOCKER_COMMANDS.md
```

#### Tests & Verification
- Docker build structure verified ✅
- Multi-stage build configured correctly ✅
- Health check endpoint configured ✅
- Dev compose with live reload ready ✅
- Documentation complete and comprehensive ✅

---

### �📦 Initial Setup & Boilerplate (March 10, 2026)

#### Added
- **Clean Project Architecture**
  - `cmd/api/main.go` - Application entry point with structured startup
  - `internal/config/` - Configuration management with environment variables
  - `internal/server/` - HTTP server lifecycle management
  - `internal/http/` - HTTP handlers and routing

- **Single Health Check Endpoint**
  - `GET /up` - Returns `{"status":"up"}` with 200 OK
  - JSON response format with proper Content-Type header
  - Structured logging for all requests

- **Comprehensive Unit Tests**
  - `internal/config/config_test.go` - Tests for config loading and env vars
  - `internal/http/router_test.go` - Tests for HTTP handlers, status codes, and content type
  - `internal/server/server_test.go` - Tests for server initialization and configuration
  - All tests passing ✅

- **Development Tools & Configuration**
  - `.air.toml` - Live reload configuration for automatic app restart on code changes
  - `Makefile` - Common development tasks (dev, run, build, test, lint, fmt, clean)
  - `dev.sh` - Convenient development script that auto-detects air installation
  - `.gitignore` - Proper ignore patterns for Go projects

- **Documentation**
  - `README.md` - Comprehensive guide covering:
    - Project features and architecture
    - Installation and setup instructions
    - Multiple ways to run the application
    - API endpoint documentation
    - Testing procedures with coverage
    - Development with live reload
    - How to add new endpoints
    - Troubleshooting guide

#### Best Practices Implemented
- ✅ Separation of Concerns (Config, Server, HTTP packages)
- ✅ Structured Logging (JSON format using `log/slog`)
- ✅ Environment-based Configuration (12-factor app)
- ✅ Graceful Server Management (timeouts, proper shutdown)
- ✅ Comprehensive Tests (config, handlers, server)
- ✅ Live Reload Development (air integration)
- ✅ Makefile for Common Tasks
- ✅ Clean Code with Documentation Comments

#### Dependencies
- Go 1.25.0
- No external dependencies (uses Go standard library only)
- Optional: `air` for live reload development

#### How to Get Started
```bash
# Clone the repository
cd /Users/tomyo/git/golang-sandbox

# Install air for live reload (optional)
go install github.com/air-verse/air@latest

# Run tests
go test ./...

# Start with live reload
./dev.sh
# Or
make dev

# Or run directly
go run cmd/api/main.go
```

#### Test Coverage
- Configuration loading and environment variable handling
- HTTP handler responses and content types
- Server initialization and address configuration
- All tests pass successfully
