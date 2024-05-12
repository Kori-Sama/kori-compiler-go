run:
	@go run ./cmd/main.go

build:
	@go build -o ./bin/c ./cmd/main.go

test:
	@go test -v ./...

.PHONY: run build test