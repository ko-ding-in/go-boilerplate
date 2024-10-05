# Golang Boilerplate Project

This is a basic Go project boilerplate with a focus on clean architecture, modularity, and testing coverage.

## Project Structure

```bash
├── cmd                 # Main applications of the project
├── internal            # Private application code
│   ├── config          # Configuration logic
│   ├── handlers        # API Handlers (Controllers)
│   ├── repository      # Database Repositories
│   ├── services        # Business logic
│   └── utils           # Utility functions
├── pkg                 # Public libraries shared across projects
├── test                # Test utilities and mock data
├── go.mod              # Go module file
└── go.sum              # Go dependencies lock file
