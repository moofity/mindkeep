# MindKeep - Project Overview

## Purpose
MindKeep is a modern web application designed to help users organize and manage their thoughts, ideas, and knowledge. The project follows a monorepo architecture using pnpm and Nx for efficient development and maintenance.

## Project Structure
The project is organized as a monorepo with the following structure:

```
mindkeep/
├── apps/                    # Applications
│   ├── server/             # Go/Gin backend
│   └── frontend/           # Angular frontend
├── docs/                    # Project documentation
└── tools/                   # Build tools and scripts
```

## Technology Stack

### Backend (Server)
- **Language**: Go
- **Framework**: Gin
- **Package Manager**: Go Modules
- **Runtime**: Go 1.21+

### Frontend (Client)
- **Framework**: Angular
- **Package Manager**: pnpm
- **Build Tool**: Nx
- **Language**: TypeScript

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

## Next Steps
- Implement the frontend application
- Set up shared libraries
- Add authentication and authorization
- Implement core features
