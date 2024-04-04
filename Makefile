.PHONY: fmt check test

all: fmt check test

fmt:
	gofumpt -s -w -l .
	@echo 'goimports' && goimports -w -local github.com/omegaatt36/film36exp $(shell find . -type f -name '*.go')
	go mod tidy

check:
	go vet -all ./...
	golangci-lint run
	misspell -error */**
	@echo 'staticcheck' && staticcheck $(shell go list ./...)

test:
	go test ./...
