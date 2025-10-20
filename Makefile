.PHONY: run build swagger test test-verbose test-coverage test-watch clean mock

# Build tags for Sonic JSON (faster than standard library)
BUILD_TAGS = -tags=sonic

run:
	go run $(BUILD_TAGS) cmd/main.go

build:
	mkdir -p bin
	go build $(BUILD_TAGS) -o bin/gin-swagger-api cmd/main.go

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

