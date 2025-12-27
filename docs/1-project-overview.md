# MindKeep - Project Overview

## Purpose
MindKeep is a modern web application designed to help users organize and manage their thoughts, ideas, and knowledge. The project follows a monorepo architecture using pnpm for efficient development and maintenance.

## Project Structure
The project is organized as a monorepo with the following structure:

```
mindkeep/
├── backend/                 # TypeScript/Koa backend server
├── cli/                      # Command-line interface tools
├── client/                   # TypeScript/Angular frontend (planned)
├── docs/                     # Project documentation
└── [root]/                   # Shared configuration and dependencies
```

## Technology Stack

### Backend
- **Language**: TypeScript 5.9.3
- **Framework**: Koa (planned)
- **Package Manager**: pnpm 10.26.2
- **Runtime**: Node.js v24.12.0
- **Type Definitions**: @types/node 25.0.3

### CLI
- **Language**: TypeScript 5.9.3
- **Package Manager**: pnpm 10.26.2
- **Runtime**: Node.js v24.12.0

### Frontend (Planned)
- **Framework**: Angular
- **Package Manager**: pnpm
- **Build Tool**: Nx (planned)
- **Language**: TypeScript

## Development Setup
The project uses:
- **pnpm** for monorepo package management with workspace support
- **TypeScript 5.9.3** for type safety across all packages
- **ESLint 9.39.2** with TypeScript ESLint for code linting
- **Prettier 3.7.4** for code formatting
- **EditorConfig** for consistent code formatting across editors
- **Corepack** support via `packageManager` field for consistent pnpm version

## Configuration

### TypeScript
- Root `tsconfig.json` provides base configuration
- Package-specific `tsconfig.json` files extend the root config
- Target: ES2024 (matching Node.js 24.12.0 capabilities)
- Module system: ESNext with bundler resolution

### Code Quality
- **ESLint**: Configured with TypeScript ESLint parser and recommended rules
- **Prettier**: Integrated with ESLint via `eslint-config-prettier`
- **EditorConfig**: Enforces consistent formatting (tabs, LF line endings)

### Workspace Configuration
- `pnpm-workspace.yaml` defines workspace packages: `server` (backend) and `cli`
- Shared dev dependencies installed at root level
- Package-specific dependencies installed per package

## Documentation Structure
This documentation is organized in multiple markdown files:
- `1-project-overview.md` (this file): General project structure and purpose
- Additional documentation files will be added with numeric prefixes for logical ordering

## Current Status
- ✅ Monorepo structure set up with pnpm workspaces
- ✅ TypeScript configuration for backend and CLI
- ✅ ESLint and Prettier configured
- ✅ Backend package initialized
- ✅ CLI package initialized
- ⏳ Frontend application (to be implemented)
- ⏳ Koa framework setup (to be implemented)
- ⏳ Nx build tool setup (planned for frontend)

## Next Steps
- Set up Koa framework in backend
- Implement backend API structure
- Create frontend application with Angular
- Set up Nx for frontend build automation
- Implement shared libraries
- Add authentication and authorization
- Implement core features
