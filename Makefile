DOCKER=docker

PROJECT_NAME=film36exp

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

setup: setup-postgres

setup-adminer:
	@if ! $(DOCKER) ps | /bin/grep ${PROJECT_NAME}-adminer-local; then \
		$(DOCKER) run --name ${PROJECT_NAME}-adminer-local \
			--link ${PROJECT_NAME}-postgres-local:postgres \
			-p 8080:8080 \
			--restart always \
			-d adminer:latest;\
	fi

setup-postgres:
	@if ! $(DOCKER) ps | /bin/grep ${PROJECT_NAME}-postgres-local; then \
		$(DOCKER) run --name ${PROJECT_NAME}-postgres-local \
			-p 5432:5432 \
			-v ${PROJECT_NAME}_data:/var/lib/postgresql/data \
			-e POSTGRES_DB=${DB_NAME} \
			-e POSTGRES_USER=${DB_USER} \
			-e POSTGRES_PASSWORD=${DB_PASSWORD} \
			--restart always \
			-d postgres:16;\
	fi

remove:
	$(DOCKER) rm -f ${PROJECT_NAME}-postgres-local

test:
	go test ./...
