# Go API

A RESTful API built with Go for managing cryptocurrency coin balance operations. This project demonstrates best practices in API development including middleware authentication, organized project structure, and database integration.

## Project Overview

This Go API provides endpoints to retrieve and manage cryptocurrency coin balances with built-in authorization middleware for secure access control.

## Project Structure

```
.
├── api/
│   └── api.go              # API initialization and setup
├── cmd/
│   └── api/
│       └── main.go         # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── api.go          # Handler setup and routing
│   │   └── get_coin_balance.go  # Coin balance retrieval handler
│   ├── middleware/
│   │   └── authorization.go    # Authorization middleware
│   └── tools/
│       ├── database.go     # Database interface and setup
│       └── mockdb.go       # Mock database for testing
├── go.mod                  # Go module definition
└── go.sum                  # Go module dependencies
```

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Abhishekkapoor98/goapi.git
cd goapi
```

2. Download dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o bin/api ./cmd/api
```

## Usage

### Running the API

```bash
# Using go run
go run ./cmd/api/main.go

# Or using the compiled binary
./bin/api
```

The API will start and listen on a configured port (check main.go for the default port).

## API Endpoints

### Get Coin Balance
- **Endpoint**: `GET /coin-balance`
- **Authorization**: Requires valid authorization header
- **Description**: Retrieves the balance of a specific cryptocurrency coin
- **Response**: Returns coin balance information

## Key Features

- **Authorization Middleware**: Secure endpoints with token-based authentication
- **Modular Architecture**: Clean separation of concerns with handlers, middleware, and tools
- **Database Integration**: Abstracted database layer with mock implementation for testing
- **RESTful Design**: Standard HTTP methods and status codes

## Project Components

### API (api/api.go)
Sets up the API server, configures routes, and applies middleware.

### Handlers (internal/handlers/)
- `api.go`: Routes HTTP requests to appropriate handlers
- `get_coin_balance.go`: Implements the coin balance retrieval logic

### Middleware (internal/middleware/)
- `authorization.go`: Validates incoming requests for proper authorization

### Tools (internal/tools/)
- `database.go`: Database interface and configuration
- `mockdb.go`: Mock implementation for testing without a real database


## Author

Abhishek Kapoor

## Support

For issues, questions, or suggestions, please open an issue on the GitHub repository.
