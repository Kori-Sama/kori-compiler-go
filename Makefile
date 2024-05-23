run:
	@go run ./cmd/main.go

build:
	@CGO_ENABLED=0 
	@GOOS=linux 
	@GOARCH=amd64
	@go build -o ./bin/koric ./cmd/main.go

build-win:
	@CGO_ENABLED=0 
	@GOOS=windows 
	@GOARCH=amd64
	@go build -o ./bin/win/koric ./cmd/main.go

build-mac:
	@CGO_ENABLED=0 
	@GOOS=darwin 
	@GOARCH=amd64
	@go build -o ./bin/mac/koric ./cmd/main.go

test:
	@go test -v ./...

.PHONY: run build test