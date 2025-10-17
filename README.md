# Meek API

A production-ready REST API template built with Go, following Clean Architecture and Hexagonal Architecture patterns.

## Features

- Clean Architecture (Hexagonal/Ports & Adapters pattern)
- Gin web framework
- BDD testing with Ginkgo v2 and Gomega
- Automated mock generation with Mockery
- Swagger/OpenAPI documentation
- RESTful API design
- Comprehensive test coverage
- Example CRUD operations for users

## Prerequisites

- Go 1.21 or higher
- Make (optional but recommended)

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd meek

# Install dependencies
go mod download

# Install required tools
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/onsi/ginkgo/v2/ginkgo@latest
go install github.com/vektra/mockery/v2@latest

# Generate mocks
make mock

# Generate Swagger documentation
make swagger
```

## Running the Application

### Using Make

```bash
# Generate Swagger docs and run the server
make dev

# Or just run the server
make run

# Build binary
make build

# Regenerate Swagger docs only
make swagger
```

### Using Go commands

```bash
# Generate Swagger documentation
swag init

# Run the server
go run main.go
```

The server will start on `http://localhost:8081`

## Testing

```bash
# Run tests with Ginkgo (recommended)
make ginkgo

# Or run Ginkgo directly
ginkgo run --randomize-all --race --cover -r

# Watch mode for continuous testing
make ginkgo-watch

# Or use standard go test
make test
```

## API Documentation

Once the server is running, you can access:

- **Swagger UI**: http://localhost:8081/swagger/index.html
- **Health Check**: http://localhost:8081/health

## API Endpoints

### Health
- `GET /health` - Health check endpoint

### Users
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create a new user
- `PUT /api/v1/users/:id` - Update a user
- `DELETE /api/v1/users/:id` - Delete a user

## Example API Calls

### Create a user
```bash
curl -X POST http://localhost:8081/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Get all users
```bash
curl http://localhost:8081/api/v1/users
```

### Get user by ID
```bash
curl http://localhost:8081/api/v1/users/1
```

### Update a user
```bash
curl -X PUT http://localhost:8081/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8081/api/v1/users/1
```

## Project Structure

```
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── domain/              # Domain entities (pure business logic, NO tags)
│   │   └── user.go
│   ├── port/                # Port interfaces (contracts)
│   │   └── service/         # Service interfaces
│   ├── service/             # Service implementations (business logic)
│   │   └── usersvc/
│   └── handler/             # HTTP handlers (adapters)
│       ├── health.go
│       └── userhdl/
├── mock/                    # Generated mocks
├── docs/                    # Generated Swagger documentation
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
├── Makefile                 # Build and run commands
├── CLAUDE.md                # Complete project documentation
└── README.md                # This file
```

## Development

### Regenerating Swagger Docs

After modifying API comments in the code, regenerate the docs:

```bash
make swagger
# or
swag init -g cmd/main.go
```

### Swagger Annotations

The API uses Swagger annotations in comments. Example:

```go
// @Summary Get all users
// @Description Get all users from the system
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(c *gin.Context) {
    // handler code
}
```

## Documentation

For complete project documentation, architecture details, testing guidelines, and development best practices, see [CLAUDE.md](CLAUDE.md).

Key topics covered:
- Architecture & Design patterns
- Complete project history
- Testing strategy with Ginkgo/Gomega
- Development workflow
- Design decisions and rationale
- Best practices and conventions

## License

MIT
