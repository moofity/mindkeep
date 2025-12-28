# MindKeep Backend

Go backend server for the MindKeep application, built with the Gin web framework and PostgreSQL.

## Architecture

The backend follows a clean, modular architecture:

```
backend/
├── cmd/api/              # Application entry point
│   └── main.go          # Server initialization and graceful shutdown
├── internal/
│   ├── database/        # Database service
│   │   ├── database.go # Connection management and health checks
│   │   └── database_test.go # Integration tests
│   └── server/          # HTTP server and routing
│       ├── server.go    # Server configuration
│       ├── routes.go    # Route registration and handlers
│       └── routes_test.go # Route tests
├── docker-compose.yml   # PostgreSQL container configuration
├── Makefile            # Build and development commands
└── go.mod              # Go module dependencies
```

## Technology Stack

- **Go 1.25.5**: Programming language
- **Gin v1.11.0**: HTTP web framework
- **PostgreSQL**: Database (via pgx/v5 driver)
- **Testcontainers**: Integration testing
- **godotenv**: Environment variable management
- **gin-contrib/cors**: CORS middleware

## Prerequisites

- Go 1.25.5 or later
- Docker or Podman (for database container)
- Make (for running development commands)

## Environment Variables

Create a `.env` file in the backend directory with the following variables:

```env
PORT=8080
BLUEPRINT_DB_HOST=localhost
BLUEPRINT_DB_PORT=5432
BLUEPRINT_DB_USERNAME=postgres
BLUEPRINT_DB_PASSWORD=postgres
BLUEPRINT_DB_DATABASE=mindkeep
BLUEPRINT_DB_SCHEMA=public
```

## Getting Started

### 1. Start the Database

Start the PostgreSQL container using Podman or Docker:

```bash
make podman-run
```

Or using Docker Compose directly:

```bash
docker-compose up -d
```

### 2. Build the Application

```bash
make build
```

This will create a `main` binary in the backend directory.

### 3. Run the Application

```bash
make run
```

The server will start on the port specified in your `.env` file (default: 8080).

### 4. Verify the Server

Check the health endpoint:

```bash
curl http://localhost:8080/health
```

## Available Make Commands

| Command | Description |
|---------|-------------|
| `make all` | Build and run tests |
| `make build` | Build the application binary |
| `make run` | Run the application |
| `make test` | Run all tests |
| `make itest` | Run integration tests |
| `make podman-run` | Start database container |
| `make podman-down` | Stop database container |
| `make watch` | Run with live reload (requires air) |
| `make clean` | Remove build artifacts |

## Development

### Live Reload

For development with automatic reloading, use:

```bash
make watch
```

This uses [Air](https://github.com/air-verse/air) for live reloading. If not installed, it will prompt you to install it.

### Running Tests

Run all tests:

```bash
make test
```

Run integration tests (requires database container):

```bash
make itest
```

### Database Management

Start the database container:

```bash
make podman-run
```

Stop the database container:

```bash
make podman-down
```

## API Endpoints

### Health Check

- **GET** `/health` - Returns database health status and connection statistics

### Root

- **GET** `/` - Returns a simple "Hello World" message

## Project Structure

### Entry Point (`cmd/api/main.go`)

Handles:
- Server initialization
- Graceful shutdown on SIGINT/SIGTERM
- 5-second timeout for in-flight requests

### Server (`internal/server/`)

- **server.go**: Server configuration and initialization
- **routes.go**: Route registration, middleware setup, and request handlers

### Database (`internal/database/`)

- **database.go**: Database connection management, health checks, and connection pooling
- **database_test.go**: Integration tests using testcontainers

## Configuration

The application uses environment variables for configuration, loaded automatically via `godotenv/autoload`. All database and server settings can be configured through the `.env` file.

## Graceful Shutdown

The server implements graceful shutdown:
- Listens for SIGINT/SIGTERM signals
- Allows 5 seconds for in-flight requests to complete
- Closes database connections cleanly

## Testing

### Unit Tests

Standard Go testing for individual components.

### Integration Tests

Integration tests use testcontainers to spin up a real PostgreSQL instance, ensuring database operations work correctly in a containerized environment.

## Contributing

1. Follow Go best practices and conventions
2. Write tests for new features
3. Ensure all tests pass before submitting
4. Use `make test` to verify everything works
