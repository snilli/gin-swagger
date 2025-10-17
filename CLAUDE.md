# CLAUDE.md - Meek API Project

> Complete project documentation and development guidelines for Claude Code and developers.

**Last Updated:** October 17, 2024
**Project:** Meek API
**Version:** 1.0.0

---

## Table of Contents

1. [Project Overview](#project-overview)
2. [Project History](#project-history)
3. [Architecture & Design](#architecture--design)
4. [Directory Structure](#directory-structure)
5. [Technology Stack](#technology-stack)
6. [Development Setup](#development-setup)
7. [Testing Guidelines](#testing-guidelines)
8. [Development Workflow](#development-workflow)
9. [API Documentation](#api-documentation)
10. [Design Decisions](#design-decisions)
11. [Best Practices](#best-practices)
12. [Future Enhancements](#future-enhancements)

---

## Project Overview

**Meek API** is a RESTful API service built with Go, following Clean Architecture principles and Hexagonal Architecture patterns. This project serves as a production-ready template demonstrating best practices in Go web development.

### Project Goals
- Build a production-ready REST API template
- Implement Clean Architecture with clear layer boundaries
- Achieve comprehensive test coverage using BDD (Ginkgo/Gomega)
- Provide automatic API documentation with Swagger/OpenAPI
- Follow Go best practices and community standards

### Key Features
- ✅ RESTful API with Gin web framework
- ✅ Clean Architecture (Hexagonal/Ports & Adapters)
- ✅ BDD testing with Ginkgo v2 and Gomega
- ✅ Automated mock generation with Mockery
- ✅ Swagger/OpenAPI documentation
- ✅ CRUD operations for User entity
- ✅ Comprehensive test coverage
- ✅ Health check endpoint

---

## Project History

### Initial Commit - "init" (f05d4d2)
**Created:**
- Basic project structure with Gin framework
- CRUD endpoints for User management
- Swagger/OpenAPI integration
- Makefile for automation
- README documentation

**Files:**
```
cmd/main.go                 # Application entry point
internal/domain/user.go     # Domain model
internal/handler/health.go  # Health check handler
internal/handler/userhdl/   # User handlers (CRUD)
internal/service/usersvc/   # User service layer (CRUD)
internal/port/service/      # Service port interfaces
docs/                       # Swagger documentation
```

### Second Commit - "fix mockery config" (9dc6c34)
**Fixed:**
- Mockery configuration for proper mock generation
- Separate mock packages for services and repositories
- Package naming conventions: `mock/mockservice/` and `mock/mockrepository/`

### Recent Development - Ginkgo/Gomega Setup
**Added:**
- Ginkgo v2 testing framework
- Gomega assertion library
- Test suite bootstrap
- Example BDD tests
- Comprehensive testing guidelines

---

## Architecture & Design

### Architecture Pattern: Clean Architecture + Hexagonal Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    External World                        │
│  (HTTP Clients, Databases, External Services)           │
└────────────────────┬────────────────────────────────────┘
                     │
        ┌────────────▼────────────┐
        │   Adapters (Handlers)   │  ◄── HTTP/gRPC handlers
        │   - userhdl/            │
        │   - health.go           │
        └────────────┬────────────┘
                     │
        ┌────────────▼────────────┐
        │    Ports (Interfaces)   │  ◄── Service interfaces
        │   - port/service/       │
        └────────────┬────────────┘
                     │
        ┌────────────▼────────────┐
        │  Business Logic (Core)  │  ◄── Service implementations
        │   - service/usersvc/    │
        └────────────┬────────────┘
                     │
        ┌────────────▼────────────┐
        │   Domain (Entities)     │  ◄── Pure business models
        │   - domain/user.go      │
        └─────────────────────────┘
```

### Layer Responsibilities

**1. Domain Layer** (`internal/domain/`)
- Pure business entities
- NO dependencies on other layers
- NO framework-specific code
- NO serialization tags (json, bson, xml, etc.)

**2. Port Layer** (`internal/port/`)
- Defines interfaces (contracts)
- Service ports: Business logic interfaces
- Repository ports: Data access interfaces
- Independent of implementation details

**3. Service Layer** (`internal/service/`)
- Implements business logic
- Implements port interfaces
- Orchestrates domain operations
- Independent of HTTP/transport concerns

**4. Handler Layer** (`internal/handler/`)
- HTTP request/response handling
- Input validation
- DTO (Data Transfer Object) models
- Maps HTTP to service calls
- Swagger annotations

---

## Directory Structure

```
meek/
├── cmd/
│   └── main.go                          # Application entry point, DI setup
│
├── internal/                            # Private application code
│   ├── domain/                          # Domain entities (pure business logic)
│   │   └── user.go                      # User entity (NO tags!)
│   │
│   ├── port/                            # Port interfaces (contracts)
│   │   ├── service/usersvc/            # Service port interfaces
│   │   │   └── service.go              # UserService interface
│   │   └── repository/                  # Repository port interfaces (future)
│   │
│   ├── service/                         # Service implementations
│   │   └── usersvc/
│   │       ├── service.go              # Service struct & constructor
│   │       ├── create_user.go          # Create logic + test
│   │       ├── get_user.go             # Get logic + test
│   │       ├── get_users.go            # List logic + test
│   │       ├── update_user.go          # Update logic + test
│   │       └── delete_user.go          # Delete logic + test
│   │
│   └── handler/                         # HTTP handlers
│       ├── health.go                    # Health check handler
│       └── userhdl/
│           ├── handler.go              # Handler struct & constructor
│           ├── models.go               # DTO models (with JSON tags)
│           ├── create_user.go          # POST /users + test
│           ├── get_user.go             # GET /users/:id + test
│           ├── get_users.go            # GET /users + test
│           ├── update_user.go          # PUT /users/:id + test
│           └── delete_user.go          # DELETE /users/:id + test
│
├── mock/                                # Generated mocks
│   ├── mockservice/                     # Service mocks
│   └── mockrepository/                  # Repository mocks (future)
│
├── docs/                                # Generated Swagger docs
│   ├── docs.go
│   └── swagger.yaml
│
├── meek_suite_test.go                   # Ginkgo test suite bootstrap
├── example_test.go                      # Example Ginkgo/Gomega tests
│
├── .mockery.yaml                        # Mockery configuration
├── Makefile                             # Build automation
├── README.md                            # User-facing documentation
├── CLAUDE.md                            # This file
├── go.mod                               # Go module dependencies
└── go.sum                               # Dependency checksums
```

---

## Technology Stack

### Core Dependencies

| Category | Technology | Version | Purpose |
|----------|-----------|---------|---------|
| **Web Framework** | [Gin](https://github.com/gin-gonic/gin) | Latest | HTTP web framework |
| **API Docs** | [Swaggo](https://github.com/swaggo/swag) | Latest | OpenAPI/Swagger generation |
| **Testing** | [Ginkgo](https://github.com/onsi/ginkgo/v2) | v2.26.0 | BDD testing framework |
| **Assertions** | [Gomega](https://github.com/onsi/gomega) | v1.38.2 | Matcher/assertion library |
| **Mocking** | [Mockery](https://github.com/vektra/mockery) | Latest | Mock generation |
| **Test Utils** | [Testify](https://github.com/stretchr/testify) | Latest | Mock assertions |

### Development Tools
- **Go:** 1.21+
- **Build Tool:** Make
- **Version Control:** Git
- **Package Manager:** Go Modules

---

## Development Setup

### Prerequisites

```bash
# Install required tools
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/onsi/ginkgo/v2/ginkgo@latest
go install github.com/vektra/mockery/v2@latest
```

### Initial Setup

```bash
# 1. Clone and enter directory
git clone <repository-url>
cd meek

# 2. Install dependencies
go mod download

# 3. Generate mocks
make mock

# 4. Generate Swagger documentation
make swagger

# 5. Run tests
ginkgo run --randomize-all --race --cover

# 6. Run the application
make dev
```

### Setup Checklist

- [ ] Go 1.21+ installed
- [ ] Swag CLI installed
- [ ] Ginkgo CLI installed
- [ ] Mockery installed
- [ ] Dependencies downloaded
- [ ] Mocks generated
- [ ] Swagger docs generated
- [ ] Tests passing
- [ ] Server runs successfully
- [ ] Swagger UI accessible at http://localhost:8081/swagger/index.html

---

## Testing Guidelines

### BDD Testing with Ginkgo/Gomega

This project uses **Ginkgo v2** and **Gomega** for BDD-style testing.

#### Test Structure Pattern

```go
package mypackage_test

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
    var (
        service usersvc.Service
        ctx     context.Context
    )

    BeforeEach(func() {
        service = usersvc.New()
        ctx = context.Background()
    })

    Describe("CreateUser", func() {
        Context("when valid data is provided", func() {
            It("should create user successfully", func() {
                user, err := service.CreateUser(ctx, "John", "john@test.com")

                Expect(err).ToNot(HaveOccurred())
                Expect(user).ToNot(BeNil())
                Expect(user.Name).To(Equal("John"))
                Expect(user.Email).To(Equal("john@test.com"))
            })
        })

        Context("when invalid data is provided", func() {
            It("should return error", func() {
                user, err := service.CreateUser(ctx, "", "")

                Expect(err).To(HaveOccurred())
                Expect(user).To(BeNil())
            })
        })
    })
})
```

#### Ginkgo Best Practices
- Use `Describe` for grouping tests by component/function
- Use `Context` for different scenarios (success, error, edge cases)
- Use `It` for individual test cases
- Use `BeforeEach` for test setup
- Use `AfterEach` for cleanup
- Always randomize tests with `--randomize-all` flag

#### Common Gomega Matchers

```go
// Equality
Expect(result).To(Equal(expected))
Expect(result).ToNot(Equal(unexpected))

// Nil/Error checking
Expect(err).ToNot(HaveOccurred())
Expect(err).To(HaveOccurred())
Expect(result).To(BeNil())
Expect(result).ToNot(BeNil())

// Error matching
Expect(err).To(MatchError(ErrNotFound))
Expect(err).To(MatchError(ContainSubstring("not found")))

// Numeric comparisons
Expect(count).To(BeNumerically(">", 0))
Expect(value).To(BeNumerically(">=", 10))

// Collections
Expect(slice).To(HaveLen(5))
Expect(slice).To(ContainElement("item"))
Expect(slice).To(ConsistOf("a", "b", "c"))
Expect(slice).To(BeEmpty())

// Boolean
Expect(result).To(BeTrue())
Expect(result).To(BeFalse())
```

### Mock Style Conventions

- **Use `EXPECT()` pattern:** Always use `mockRepo.EXPECT().MethodName(...)` instead of `On("MethodName", ...)`
- **Avoid generic matchers:** Do NOT use `mock.Anything`, `mock.AnythingOfType()`, or `mock.MatchedBy()`
- **Use specific values:** Replace matchers with actual values:
  - Context: `context.Background()` or use `Any()` from Gomega
  - Strings: Actual values like `"user123"`, `"john@example.com"`
  - Numbers: Specific values like `int32(5)`, `int64(100)`
  - Structs: Concrete instances with real field values
- **Limited `mock.Anything`:** Only for complex objects generated by business logic

```go
// ✅ Good - Using EXPECT() with specific values
mockService.EXPECT().GetUser(context.Background(), "user-123").Return(user, nil)

// ❌ Bad - Using On() with generic matchers
mockService.On("GetUser", mock.Anything, "user-123").Return(user, nil)
```

### Running Tests

```bash
# Ginkgo with Make (recommended)
make ginkgo                    # Run all tests recursively
make ginkgo-watch              # Watch mode

# Ginkgo directly
ginkgo run --randomize-all --race --cover -r
ginkgo run --focus="CreateUser"     # Focus on specific tests
ginkgo watch -r                      # Watch mode

# Standard go test
make test
go test -v -cover ./...
```

### Test Coverage Goals
- **Target:** 80%+ code coverage
- **Current:** 100% coverage (22 specs across 3 suites)
- **Future:** Integration tests, benchmark tests

---

## Development Workflow

### Standard Development Cycle

```bash
# 1. Make code changes in internal/
# 2. Update corresponding *_test.go files

# 3. Generate mocks if interfaces changed
make mock

# 4. Run tests
ginkgo run --randomize-all --race --cover

# 5. Update Swagger docs if API changed
make swagger

# 6. Run the application
make dev

# 7. Test via Swagger UI or curl
curl http://localhost:8081/api/v1/users
```

### Adding New Features

**Example: Adding a new "Profile" entity**

```bash
# 1. Create domain entity
touch internal/domain/profile.go

# 2. Create port interface
mkdir -p internal/port/service/profilesvc
touch internal/port/service/profilesvc/service.go

# 3. Create service implementation
mkdir -p internal/service/profilesvc
touch internal/service/profilesvc/service.go
touch internal/service/profilesvc/create_profile.go
# ... other operations

# 4. Create handlers
mkdir -p internal/handler/profilehdl
touch internal/handler/profilehdl/handler.go
touch internal/handler/profilehdl/models.go
touch internal/handler/profilehdl/create_profile.go
# ... other endpoints

# 5. Write tests
touch internal/service/profilesvc/create_profile_test.go
touch internal/handler/profilehdl/create_profile_test.go

# 6. Update cmd/main.go
# - Wire up new service
# - Register new routes

# 7. Generate mocks and docs
make mock
make swagger
```

### Makefile Commands

```bash
make run           # Run the application
make build         # Build production binary
make swagger       # Generate Swagger documentation
make mock          # Generate mocks
make test          # Run standard Go tests
make ginkgo        # Run Ginkgo tests (recommended)
make ginkgo-watch  # Run Ginkgo in watch mode
make clean         # Clean build artifacts
make dev           # Generate docs and run (dev mode)
```

---

## API Documentation

### Endpoints

| Method | Endpoint | Description | Request | Response |
|--------|----------|-------------|---------|----------|
| GET | `/health` | Health check | - | `200 OK` |
| GET | `/api/v1/users` | Get all users | - | `200 OK` with array |
| GET | `/api/v1/users/:id` | Get user by ID | - | `200 OK` with object |
| POST | `/api/v1/users` | Create user | `{name, email}` | `201 Created` |
| PUT | `/api/v1/users/:id` | Update user | `{name, email}` | `200 OK` |
| DELETE | `/api/v1/users/:id` | Delete user | - | `204 No Content` |

### Swagger UI

**Access:** http://localhost:8081/swagger/index.html

**Features:**
- Interactive API testing
- Request/Response schemas
- Example values
- Try-it-out functionality

### Regenerating Documentation

```bash
# After modifying API comments
make swagger

# Or directly
swag init -g cmd/main.go
```

---

## Design Decisions

### 1. Why Clean Architecture?

**Decision:** Use Clean Architecture with Hexagonal pattern

**Rationale:**
- **Testability:** Business logic independent of frameworks
- **Flexibility:** Easy to swap implementations
- **Maintainability:** Clear separation of concerns
- **Scalability:** Well-organized code growth

**Trade-offs:**
- More initial boilerplate
- Steeper learning curve
- More abstractions

### 2. Why Ginkgo/Gomega?

**Decision:** BDD-style testing instead of standard Go tests

**Rationale:**
- **Readability:** Tests read like specifications
- **Organization:** Better structure with Describe/Context/It
- **Expressiveness:** Rich matcher library
- **Features:** Parallel testing, watch mode, focus/skip
- **Output:** Beautiful test reports

**Trade-offs:**
- Additional dependency
- Different from standard conventions
- Team learning curve

### 3. Why In-Memory Storage?

**Decision:** Use in-memory map (currently)

**Rationale:**
- **Simplicity:** Quick prototype
- **Testing:** No database setup needed
- **Learning:** Focus on architecture

**Future:** Replace with real database

### 4. Why Separate Files Per Operation?

**Decision:** One file per operation (create_user.go, get_user.go, etc.)

**Rationale:**
- **Maintainability:** Smaller, focused files
- **Git conflicts:** Less merge conflicts
- **Testing:** One-to-one mapping
- **Navigation:** Easier to find

**Trade-offs:**
- More files in project
- Some prefer single file

### 5. Why Mockery with EXPECT()?

**Decision:** Use EXPECT() pattern instead of On()

**Rationale:**
- **Type Safety:** Compile-time checking
- **Clarity:** Clear expectations
- **Maintainability:** Easier to understand
- **Best Practice:** Recommended pattern

---

## Best Practices

### Code Style

- Follow existing patterns in the codebase
- Use `gofmt` for formatting
- Follow Go naming conventions
- Write tests for all functionality
- Update Swagger comments for API changes

### Domain Layer Purity (CRITICAL)

```go
// ✅ CORRECT: Domain entities must be pure
type User struct {
    ID    string
    Name  string
    Email string
    // NO tags: json, bson, xml, etc.
}

// ❌ WRONG: Domain with transport concerns
type User struct {
    ID    string `json:"id" bson:"_id"`  // DON'T DO THIS
    Name  string `json:"name"`           // VIOLATES CLEAN ARCHITECTURE
}
```

**Rule:** Domain entities NEVER have serialization tags

### Test Organization

- Group related tests using `Describe` blocks
- Use descriptive `Context` names
- Keep `It` descriptions clear
- Always test success and error paths
- Test edge cases

### Developer Notes

- Always run `make mock` after interface changes
- Always run `make swagger` after API changes
- Use `ginkgo watch` during test development
- Check Swagger UI to verify documentation
- Follow existing patterns when adding features

---

## Future Enhancements

### Short-term
- [ ] Database integration (PostgreSQL/MongoDB)
- [ ] Repository layer implementation
- [ ] Authentication/Authorization (JWT)
- [ ] Request validation middleware
- [ ] Pagination for GetUsers
- [ ] Logging middleware
- [ ] Error handling middleware
- [ ] Environment-based configuration

### Medium-term
- [ ] Integration tests
- [ ] CORS middleware
- [ ] Rate limiting
- [ ] Soft delete
- [ ] Search/filter functionality
- [ ] Metrics (Prometheus)
- [ ] Docker containerization
- [ ] CI/CD pipeline

### Long-term
- [ ] gRPC API support
- [ ] GraphQL API
- [ ] Event-driven architecture
- [ ] Microservices split
- [ ] API versioning
- [ ] Multi-tenancy
- [ ] Caching layer (Redis)
- [ ] Message queue integration

---

## Quick Reference

### Common Commands

```bash
# Development
make dev           # Start dev server
make ginkgo        # Run tests with Ginkgo
make ginkgo-watch  # Run tests in watch mode
make mock          # Generate mocks
make swagger       # Update API docs
make build         # Build binary
make clean         # Clean artifacts
```

### Important Files

| File | Purpose |
|------|---------|
| `cmd/main.go` | Entry point, routing |
| `.mockery.yaml` | Mock configuration |
| `CLAUDE.md` | This file |
| `Makefile` | Build commands |

### Port Configuration

- **Application:** http://localhost:8081
- **Swagger UI:** http://localhost:8081/swagger/index.html
- **Health Check:** http://localhost:8081/health

---

## Project Metrics

**Code Stats:**
- Domain: ~10 lines
- Ports: ~20 lines
- Services: ~200 lines
- Handlers: ~300 lines
- Tests: ~400 lines
- **Total:** ~930 lines

**Coverage:**
- Service layer: High
- Handler layer: High
- **Overall:** 80%+ (estimated)

**Files:**
- Go files: ~30
- Test files: ~15
- Config: ~5
- Docs: ~4

---

**Maintained by:** Project Team
**Last Updated:** October 17, 2024
**Repository:** [Repository URL]
