.PHONY: build
build:
	go build -o bin/apiserver -v ./cmd/apiserver

.PHONY: test
test:
	go test -cover -race -timeout 30s ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out -coverpkg=./... -cover ./...

.DEFAULT_GOAL := build