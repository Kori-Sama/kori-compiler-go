run:
	@go run ./cmd/main.go

build:
	@go build -o ./bin/koric ./cmd/main.go

test:
	@go test -v ./...

.PHONY: run build test