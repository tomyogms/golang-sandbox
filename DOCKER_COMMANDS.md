# Docker Quick Reference

## Start with Docker Compose

### Production Build & Run
```bash
# Build and start in foreground (see logs)
docker-compose up --build

# Build and start in background
docker-compose up -d --build

# Rebuild images (skip cache)
docker-compose up --build --no-cache
```

### Development Build & Run (with Live Reload)
```bash
# Start with live reload
docker-compose -f docker-compose.dev.yml up --build

# Start in background
docker-compose -f docker-compose.dev.yml up -d --build
```

### Stop Services
```bash
# Stop production services
docker-compose down

# Stop development services
docker-compose -f docker-compose.dev.yml down
```

### View Logs
```bash
# Watch production logs
docker-compose logs -f api

# Watch development logs
docker-compose -f docker-compose.dev.yml logs -f api

# View specific lines
docker-compose logs --tail=50 api
```

## Test the API

### From Host Machine
```bash
# Health check
curl http://localhost:8080/up

# With verbose output
curl -v http://localhost:8080/up

# JSON formatted response
curl http://localhost:8080/up | jq .
```

### From Inside Container
```bash
# Execute command in running container
docker-compose exec api wget -O - http://localhost:8080/up

# Interactive shell
docker-compose exec api sh

# Run arbitrary commands
docker-compose exec api ls -la /app
```

## Manual Docker Commands

### Build Image
```bash
# Standard build
docker build -t golang-sandbox:latest .

# Build with specific tag
docker build -t golang-sandbox:v1.0.0 .

# Build without cache
docker build --no-cache -t golang-sandbox:latest .
```

### Run Container
```bash
# Run with default port
docker run -p 8080:8080 golang-sandbox:latest

# Run with custom port
docker run -p 3000:8080 golang-sandbox:latest

# Run with environment variables
docker run -p 8080:8080 -e PORT=8080 -e ENV=production golang-sandbox:latest

# Run interactively (for debugging)
docker run -it golang-sandbox:latest sh

# Run in background
docker run -d -p 8080:8080 --name my-api golang-sandbox:latest
```

### Manage Containers
```bash
# List running containers
docker ps

# List all containers
docker ps -a

# View container logs
docker logs <container-id>
docker logs golang-sandbox-api

# Follow logs in real time
docker logs -f <container-id>

# Stop container
docker stop <container-id>

# Remove container
docker rm <container-id>

# Remove all stopped containers
docker container prune
```

### Inspect Images
```bash
# List images
docker images

# Image details
docker inspect golang-sandbox:latest

# View image layers
docker history golang-sandbox:latest

# Remove image
docker rmi golang-sandbox:latest
```

## Debugging

### Check Container Status
```bash
# Health status
docker-compose ps

# Detailed info
docker inspect golang-sandbox-api
```

### Access Running Container
```bash
# Open shell in container
docker-compose exec api sh

# Run Go commands
docker-compose exec api go version

# Check environment variables
docker-compose exec api env
```

### Check Image Contents
```bash
# List files in image
docker run golang-sandbox:latest ls -la /app

# Check binary location
docker run golang-sandbox:latest which api

# View final image size
docker images | grep golang-sandbox
```

## Performance & Optimization

### View Build Time
```bash
docker build --progress=plain -t golang-sandbox:latest .
```

### Check Image Size
```bash
# Human-readable size
docker images --format "table {{.Repository}}\t{{.Size}}" | grep golang-sandbox

# Detailed breakdown
docker history --human golang-sandbox:latest
```

### Compare Image Layers
```bash
docker inspect golang-sandbox:latest | grep -A 100 "GraphDriver"
```

## Common Issues & Solutions

### Port Already in Use
```bash
# Change port in docker-compose.yml
# ports:
#   - "3000:8080"

# Or kill process using port 8080
lsof -i :8080
kill -9 <pid>
```

### Container Exits Immediately
```bash
# Check logs
docker logs golang-sandbox-api

# Run with interactive shell to debug
docker run -it golang-sandbox:latest sh
```

### Cannot Connect to API
```bash
# Verify container is running
docker ps

# Check if port is exposed
docker port golang-sandbox-api

# Test from inside container
docker-compose exec api wget -O - http://localhost:8080/up
```

### Build Cache Issues
```bash
# Clear build cache
docker builder prune

# Rebuild without cache
docker-compose up --build --no-cache
```

## Clean Up

### Remove Unused Resources
```bash
# Remove stopped containers
docker container prune

# Remove dangling images
docker image prune

# Remove unused volumes
docker volume prune

# Remove everything (be careful!)
docker system prune -a
```

## Environment Setup

### Add to .env (optional)
```bash
PORT=8080
ENV=production
```

### Load from .env file
```bash
docker-compose --env-file .env up --build
```
