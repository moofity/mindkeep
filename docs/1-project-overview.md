# Project Overview

## Purpose
MindKeep is a modern web application designed to help users organize and manage their thoughts, ideas, and knowledge. The project follows a monorepo architecture using pnpm and Nx for efficient development and maintenance.

## Project Structure
The project is organized as a monorepo with the following structure:

```
mindkeep/
├── server/                  # Go/Gin backend
├── client/                  # Typescript/Angular frontent
├── docs/                    # Project documentation
└── tools/                   # Build tools and scripts
```

## Technology Stack

### Backend (Server)
- **Language**: Go
- **Framework**: Gin
- **Package Manager**: Go Modules
- **Runtime**: Go 1.21+
- **Architecture**: Monolithic application with modular design
- **Modules**: The backend is organized into modules, each responsible for a specific domain of functionality and its associated business logic:
  - **social-graph**: Core module for managing relationships and connections (described in detail in the [Social graph module documentation](2-social-graph-module.md))
  - *Additional modules will be added as the project evolves*

### Frontend (Client)
- **Language**: TypeScript
- **Framework**: Angular
- **Package Manager**: pnpm
- **Build Tool**: Nx

## Infrastructure
- **Graph Database**: Neo4j for storing and managing relationship data
- *Additional infrastructure components will be added as the project evolves*

## Development Setup
The project uses:
- pnpm for frontend package management
- Go Modules for backend dependency management
- Nx for frontend build automation
- TypeScript for frontend type safety
- EditorConfig for consistent code formatting

## Documentation Structure
This documentation is organized in multiple markdown files:
- `1-project-overview.md` (this file): General project structure and purpose
- Additional documentation files will be added with numeric prefixes for logical ordering
