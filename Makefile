.PHONY: run build swagger test mock clean

run:
	go run cmd/main.go

build:
	go build -o bin/meek cmd/main.go

swagger:
	swag init -g cmd/main.go

mock:
	mockery && sed -i '' 's/^package port$$/package mock/' mock/*.go

test:
	go test -v -cover ./...

clean:
	rm -rf bin/ docs/ mock/

dev: swagger run
