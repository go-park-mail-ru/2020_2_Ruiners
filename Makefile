.PHONY: build
build:
	go build -o bin/apiserver -v ./cmd/apiserver
	go build -o bin/session -v ./internal/pkg/microsevice/session/cmd
	go build -o bin/auth -v ./internal/pkg/microsevice/auth/cmd
	go build -o bin/rate -v ./internal/pkg/microsevice/rate/cmd

.PHONY: test
test:
	go test -cover -race -timeout 30s ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out -coverpkg=./... -cover ./...

.DEFAULT_GOAL := build