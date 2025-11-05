.PHONY: run run-graphql build build-graphql build-all swagger test test-verbose test-coverage test-watch clean mock orm local-orm help

# Build tags for Sonic JSON (faster than standard library)
BUILD_TAGS = -tags=sonic

# ORM Provider settings
ORM_REMOTE = github.com/snilli/ormprovider@latest
ORM_LOCAL = ../ormprovider

# Run REST API server
run:
	go run $(BUILD_TAGS) cmd/api/main.go

# Run GraphQL server
run-graphql:
	go run $(BUILD_TAGS) cmd/graphql/main.go

# Build GraphQL server
build-graphql:
	mkdir -p bin
	go build $(BUILD_TAGS) -o bin/graphql-server cmd/graphql/main.go

# Build both servers
build-all: build build-graphql

swagger:
	swag init -g cmd/main.go

mock:
	mockery

test:
	ginkgo run --randomize-all --cover -r

test-verbose:
	ginkgo run --randomize-all --cover -r -v

test-coverage:
	ginkgo run --randomize-all --cover -r --cover-profile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

test-watch:
	ginkgo watch -r

clean:
	rm -rf bin/ docs/ coverage.out coverage.html

orm:
	go mod edit -dropreplace github.com/snilli/ormprovider || true
	go get $(ORM_REMOTE)
	go mod tidy

local-orm:
	go mod edit -replace github.com/snilli/ormprovider=$(ORM_LOCAL)
	go mod tidy

