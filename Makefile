.PHONY: run build swagger test ginkgo ginkgo-watch mock clean dev

run:
	go run cmd/main.go

build:
	go build -o bin/meek cmd/main.go

swagger:
	swag init -g cmd/main.go

mock:
	mockery

test:
	go test -v -cover ./...

ginkgo:
	ginkgo run --randomize-all --race --cover -r

ginkgo-watch:
	ginkgo watch -r

clean:
	rm -rf bin/ docs/ mock/

dev: swagger run
