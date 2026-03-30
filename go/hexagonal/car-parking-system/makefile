.PHONY: build run test lint vet fmt

build-api:
	go build -o bin/api ./cmd/api

run:
	go run ./cmd/api

test:
	go test ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...