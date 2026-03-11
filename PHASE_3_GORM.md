# Phase 3: PostgreSQL with GORM - Setup Complete ✅

## Overview

Your Golang Sandbox application now has a fully configured PostgreSQL database layer using GORM (Go Object-Relational Mapping) with best practices for production-ready code.

## Key Features Implemented

✅ **GORM Integration**
- Modern Go ORM for type-safe database operations
- Supports complex queries and relationships
- Built-in transaction support
- Automatic migrations capability

✅ **PostgreSQL Container**
- PostgreSQL 17-alpine for lightweight deployment
- Persistent data volumes for both dev and prod
- Health checks for reliable startup sequencing
- Proper connection pooling configuration

✅ **Production-Ready Architecture**
- Connection pooling with configurable limits
- Graceful connection lifecycle management
- Structured logging integration
- Environment-based configuration

✅ **Good Practices**
- Dependency injection pattern
- Separation of concerns (config, database, server layers)
- Error handling with context
- Resource cleanup in shutdown

## Files Modified (5)

### 1. **go.mod**
```diff
+ gorm.io/gorm v1.31.1
+ gorm.io/driver/postgres v1.6.0
+ github.com/jackc/pgx/v5 v5.6.0
+ (and related dependencies)
```

### 2. **internal/db/db.go**
- Complete rewrite to use GORM instead of raw SQL driver
- Added `Database` struct wrapping `*gorm.DB`
- Added `Config` struct for connection parameters
- Connection pooling: max 100 open connections, 10 idle
- Proper error handling and resource cleanup
- Structured logger integration

**Key Functions:**
- `New(cfg Config) (*Database, error)` - Creates GORM connection
- `Close() error` - Gracefully closes database connection

### 3. **internal/config/config.go**
- Extended `Config` struct with `Database db.Config` field
- Added database configuration from environment variables
- Defaults: localhost, port 5432, postgres/postgres user
- Validation for required database host

**Environment Variables:**
- `DB_HOST` - Database host (default: localhost)
- `DB_PORT` - Database port (default: 5432)
- `DB_USER` - Database user (default: postgres)
- `DB_PASSWORD` - Database password (default: postgres)
- `DB_NAME` - Database name (default: golang_sandbox)
- `DB_SSLMODE` - SSL mode (default: disable)

### 4. **internal/server/server.go**
- Added `database *db.Database` field to Server struct
- Updated `Start()` to initialize database before listening
- Deferred database cleanup during server lifetime
- Enhanced `Shutdown()` to close database connection
- Added database connection logging

### 5. **docker-compose.yml (Production)**
```yaml
services:
  postgres:
    image: postgres:17-alpine
    ports: 5432:5432
    volumes: postgres_data:/var/lib/postgresql/data
    healthcheck: pg_isready -U postgres
  
  api:
    environment:
      - DB_HOST=postgres
      - DB_PORT=5432
      - DB_USER=postgres
      - DB_PASSWORD=postgres
      - DB_NAME=golang_sandbox
    depends_on:
      postgres:
        condition: service_healthy
```

### 6. **docker-compose.dev.yml (Development)**
- Identical PostgreSQL setup with separate dev volume
- API service with volume mounts for hot-reload
- same environment variables as production
- Supports development workflow with `air` auto-reload

## Architecture

```
┌─────────────────────────────────────────────┐
│         main.go                             │
│  (Loads config, creates server)             │
└────────────────┬────────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────────┐
│  Server (internal/server)                   │
│  - Manages HTTP + Database lifecycle        │
│  - Initializes DB on Start()                │
│  - Closes DB on Shutdown()                  │
└────────────────┬────────────────────────────┘
                 │
                 ▼
┌─────────────────────────────────────────────┐
│  Database (internal/db)                     │
│  - GORM Connection                          │
│  - Connection pooling                       │
│  - Resource cleanup                         │
└────────────────┬────────────────────────────┘
                 │
                 ▼
       ┌─────────────────────┐
       │  PostgreSQL (Docker)│
       │  - Port 5432        │
       │  - Data persistence │
       └─────────────────────┘
```

## Connection Pool Configuration

```go
MaxIdleConns:    10
MaxOpenConns:    100
ConnMaxLifetime: 1 hour
```

These settings optimize for application patterns:
- **MaxIdleConns (10)**: Maintains ready connections for quick queries
- **MaxOpenConns (100)**: Handles concurrent requests adequately
- **ConnMaxLifetime (1 hour)**: Prevents connection staleness
- Total pool size adapts from idle to open based on load

## Environment Setup

### Development
```bash
# Start dev stack with auto-reload
docker compose -f docker-compose.dev.yml up --build

# View logs
docker compose -f docker-compose.dev.yml logs -f api

# Stop and cleanup
docker compose -f docker-compose.dev.yml down -v
```

### Production
```bash
# Build and start production stack
docker compose up --build

# Check database health
docker compose exec postgres pg_isready -U postgres
```

## Health Checks

### Database (PostgreSQL)
- Command: `pg_isready -U postgres`
- Interval: 10 seconds
- Timeout: 5 seconds
- Retries: 5 attempts
- Status: Used to sequence container startup

### API
- Command: `wget --quiet --tries=1 --spider http://localhost:8080/up`
- Interval: 10 seconds
- Timeout: 3 seconds
- Retries: 3 attempts
- Start Period: 5 seconds (waits before first check)

## Best Practices Implemented

### 1. **Dependency Injection**
- Config passed through layers
- Database created in Server, not globals
- Easy to test and mock

### 2. **Resource Management**
- Proper cleanup in `defer` blocks
- Database connection closed on shutdown
- No connection leaks

### 3. **Error Handling**
- Wrapped errors with context using `%w`
- Structured logging with slog
- Clear error messages for debugging

### 4. **Configuration Management**
- Environment variables for flexibility
- Defaults for development convenience
- Validation for required fields

### 5. **Logging**
- GORM logger integrated
- Structured JSON logging via slog
- Database operations logged at INFO level

### 6. **Connection Pooling**
- Connection reuse via pooling
- Limits prevent resource exhaustion
- Timeouts prevent hanging connections

## Testing Notes

The database layer is ready for repository pattern implementation:

```go
// Example: Next step - create a repository interface
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id string) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id string) error
}

// Implement with GORM
type userRepository struct {
    db *db.Database
}

func (r *userRepository) Create(ctx context.Context, user *User) error {
    return r.db.WithContext(ctx).Create(user).Error
}
```

## Next Steps

### Phase 4 Recommendations:
1. **Define Models** - Create GORM models for your domain entities
2. **Migrations** - Implement GORM auto-migrations
3. **Repositories** - Create repository pattern implementations
4. **Error Handling** - Custom error types for business logic
5. **Testing** - Repository tests with in-memory SQLite for unit tests

### Example Model Structure:
```go
// internal/models/user.go
type User struct {
    ID        string `gorm:"primaryKey"`
    Email     string `gorm:"uniqueIndex"`
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Example Repository:
```go
// internal/repository/user_repo.go
type UserRepository struct {
    db *db.Database
}

func (r *UserRepository) Create(ctx context.Context, u *User) error {
    return r.db.WithContext(ctx).Create(u).Error
}
```

## Troubleshooting

### Connection Refused
```bash
# Check if postgres is running
docker compose -f docker-compose.dev.yml ps

# View postgres logs
docker compose -f docker-compose.dev.yml logs postgres

# Restart postgres
docker compose -f docker-compose.dev.yml restart postgres
```

### Database Connection Timeout
- Verify `depends_on` with `service_healthy` condition
- Check health check thresholds if containers starting slower
- View logs for SSL/authentication errors

### Port Already in Use
```bash
# Find process using port 5432
lsof -i :5432

# Kill process
kill -9 <PID>
```

## API Endpoints

### 1. **GET /up** - Simple Health Check
Basic liveness probe (always returns 200 OK when server is running)

```bash
curl http://localhost:8080/up
```

Response:
```json
{"status":"up"}
```

### 2. **GET /health** - Detailed Health Check with Database Status
Comprehensive health check including database connection status

```bash
curl http://localhost:8080/health
```

Response (when healthy):
```json
{"status":"healthy","database":"healthy"}
```

Response (when database is unavailable):
```json
{"status":"unhealthy","database":"unhealthy"}
```

Status codes:
- `200 OK` - All services healthy
- `503 Service Unavailable` - One or more services unhealthy (database connection failed)
