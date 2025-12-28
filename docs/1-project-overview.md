# MindKeep - Project Overview

## Purpose
MindKeep is a modern web application designed to help users organize and manage their thoughts, ideas, and knowledge.

## Project Structure
The project is organized with the following structure:

```
mindkeep/
├── backend/                 # Go backend server
│   ├── cmd/api/             # API server entry point
│   ├── internal/            # Internal application code
│   │   ├── database/        # Database service and connection management
│   │   └── server/          # HTTP server, routes, and handlers
│   ├── docker-compose.yml   # Database container configuration
│   ├── Makefile             # Build and development commands
│   └── README.md            # Backend-specific documentation
├── cli/                     # Go command-line interface tools
└── docs/                    # Project documentation
```

## Technology Stack

### Backend
- **Language**: Go 1.25.5
- **Framework**: Gin (HTTP web framework)
- **Database**: PostgreSQL (using pgx/v5 driver)
- **Configuration**: Environment variables (godotenv)
- **Testing**: Testcontainers for integration tests
- **CORS**: gin-contrib/cors middleware

### CLI
- **Language**: Go

## Development Setup
The project uses:
- **Go modules** for dependency management
- **gofmt** for code formatting
- **golangci-lint** for code linting (recommended)
- **Make** for common development tasks
- **Docker/Podman** for database containers

## Backend Architecture

The backend follows a clean architecture pattern:

- **Entry Point** (`cmd/api/main.go`): Application entry point with graceful shutdown
- **Server** (`internal/server/`): HTTP server setup, route registration, and request handlers
- **Database** (`internal/database/`): Database connection management and health checks

### Key Features
- Graceful shutdown handling
- Database health monitoring
- CORS support for frontend integration
- Environment-based configuration
- Integration test support with testcontainers

## Documentation Structure
This documentation is organized in multiple markdown files:
- `1-project-overview.md` (this file): General project structure and purpose
- Additional documentation files will be added with numeric prefixes for logical ordering

## Current Status
- ✅ Backend structure implemented
- ✅ Gin framework setup
- ✅ Database service with PostgreSQL
- ✅ Health check endpoints
- ✅ CORS middleware configured

## Next Steps
- ⏳ Implement business modules (social-graph, gamify)
- ⏳ Add authentication and authorization
- ⏳ Implement API endpoints for business logic
