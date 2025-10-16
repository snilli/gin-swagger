# Meek API

A simple REST API built with Gin and Swagger.

## Features

- Gin web framework
- Swagger/OpenAPI documentation
- RESTful API design
- Example CRUD operations for users

## Prerequisites

- Go 1.21 or higher
- Make (optional)

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd meek

# Install dependencies
go mod download

# Install swag CLI (for Swagger generation)
go install github.com/swaggo/swag/cmd/swag@latest
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

The server will start on `http://localhost:8080`

## API Documentation

Once the server is running, you can access:

- **Swagger UI**: http://localhost:8080/swagger/index.html
- **Health Check**: http://localhost:8080/health

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
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### Get all users
```bash
curl http://localhost:8080/api/v1/users
```

### Get user by ID
```bash
curl http://localhost:8080/api/v1/users/1
```

### Update a user
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "email": "jane@example.com"
  }'
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## Project Structure

```
.
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── domain/              # Domain models (pure business logic)
│   │   └── user.go
│   └── handler/             # HTTP handlers
│       ├── health.go
│       └── user.go
├── docs/                    # Generated Swagger documentation
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
├── Makefile                 # Build and run commands
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

## License

MIT
