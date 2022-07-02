all: modules test build lint

modules:
	go mod tidy

test:
	go test ./...

build: build_dir
	go build -o bin/gore.exe cmd/gore/main.go

build_dir:
	mkdir -p bin

lint:
	golangci-lint run
