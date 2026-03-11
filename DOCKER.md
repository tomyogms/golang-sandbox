# Docker & Docker Compose Setup

This guide covers how to run the Golang Sandbox application using Docker and Docker Compose.

## Prerequisites

- Docker installed ([Download](https://www.docker.com/products/docker-desktop))
- Docker Compose (included with Docker Desktop)

## Quick Start with Docker Compose

### Production Build

Start the application in production mode:

```bash
docker-compose up --build
```

Or run in the background:

```bash
docker-compose up -d --build
```

The API will be available at `http://localhost:8080`

### Test the Health Endpoint

```bash
curl http://localhost:8080/up
```

Response:
```json
{"status":"up"}
```

### Stop the Services

```bash
docker-compose down
```

### View Logs

```bash
docker-compose logs -f api
```

## Development with Docker Compose

For development with live reload inside Docker:

```bash
docker-compose -f docker-compose.dev.yml up --build
```

This will:
- Mount your code as a volume
- Install `air` for live reload
- Restart the app automatically on code changes
- Keep you in the development environment

### Stop Development Services

```bash
docker-compose -f docker-compose.dev.yml down
```

### Dev Logs with Live Reload

```bash
docker-compose -f docker-compose.dev.yml logs -f api
```

## Building the Docker Image Manually

Build the image:

```bash
docker build -t golang-sandbox:latest .
```

Run a container:

```bash
docker run -p 8080:8080 golang-sandbox:latest
```

With environment variables:

```bash
docker run -p 8080:8080 -e PORT=3000 -e ENV=production golang-sandbox:latest
```

## Multi-Stage Build Details

The `Dockerfile` uses a two-stage build process:

### Stage 1: Builder
- Uses `golang:1.25-alpine` image
- Downloads dependencies
- Compiles the Go binary
- Result is a compiled binary with minimal size

### Stage 2: Runtime
- Uses slim `alpine:3.20` image
- Creates a non-root user for security
- Copies only the compiled binary
- Exposes port 8080
- Final image size: ~15-20 MB

## Docker Compose Services

### Production Compose (docker-compose.yml)

```yaml
services:
  api:
    - Builds and runs the production binary
    - Port: 8080
    - Healthcheck: Monitors /up endpoint
    - Auto-restart policy
```

### Development Compose (docker-compose.dev.yml)

```yaml
services:
  api:
    - Uses builder stage
    - Volume mounts for live code editing
    - Air installed for auto-reload
    - Development environment
```

## Environment Variables in Docker

### Available Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `ENV` | `production` | Environment name |

### Set Environment Variables

In `docker-compose.yml`:
```yaml
environment:
  - PORT=8080
  - ENV=production
```

Or with `docker run`:
```bash
docker run -e PORT=3000 -e ENV=development golang-sandbox:latest
```

## Debugging

### Check Image Contents

```bash
docker run -it golang-sandbox:latest sh
```

### Check Running Container Logs

```bash
docker logs <container-id>
docker logs golang-sandbox-api
```

### Inspect Running Process

```bash
docker exec golang-sandbox-api ps aux
```

### Test Health Endpoint in Container

```bash
docker exec golang-sandbox-api wget -O - http://localhost:8080/up
```

## Performance & Best Practices

1. **Multi-stage builds** - Reduces final image size
2. **Alpine images** - Minimal base images for security and speed
3. **Non-root user** - Container runs as `appuser` for security
4. **Healthchecks** - Monitors container health automatically
5. **.dockerignore** - Excludes unnecessary files from build context

## Pushing to Registry

### Docker Hub

```bash
# Tag the image
docker tag golang-sandbox:latest yourusername/golang-sandbox:latest

# Login (if needed)
docker login

# Push
docker push yourusername/golang-sandbox:latest
```

### Private Registry

```bash
docker tag golang-sandbox:latest registry.example.com/golang-sandbox:latest
docker push registry.example.com/golang-sandbox:latest
```

## Troubleshooting

### Port Already in Use

Change the port in docker-compose.yml:
```yaml
ports:
  - "3000:8080"  # External:Internal
```

### Container Exits Immediately

Check logs:
```bash
docker-compose logs api
```

### Cannot Connect to API

Verify the container is running:
```bash
docker-compose ps
```

Test from inside container:
```bash
docker-compose exec api wget -O - http://localhost:8080/up
```
