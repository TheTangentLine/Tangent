BIN     := bin/tangent
PKG     := ./cmd/tangent
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)

# Load .env (if present) and export it to commands run by make.
-include .env
export

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@grep -hE '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the tangent binary into bin/
	go build -trimpath -ldflags "-s -w -X main.version=$(VERSION)" -o $(BIN) $(PKG)

.PHONY: run-api run-worker run-migrate
run-api: ## Run the API locally (reads .env)
	go run $(PKG) api
run-worker: ## Run the worker locally (reads .env)
	go run $(PKG) worker
run-migrate: ## Run migrations locally (reads .env)
	go run $(PKG) migrate

.PHONY: test
test: ## Run unit tests with the race detector
	go test -race -count=1 ./...

.PHONY: cover
cover: ## Run tests and open the coverage report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: lint
lint: ## Run golangci-lint
	@command -v golangci-lint >/dev/null || { echo "golangci-lint not found: brew install golangci-lint"; exit 1; }
	golangci-lint run ./...

.PHONY: fmt
fmt: ## Format code
	gofmt -w -s .

.PHONY: tidy
tidy: ## Tidy go.mod and fail if it changed (used in CI)
	go mod tidy
	git diff --exit-code go.mod $(wildcard go.sum)

.PHONY: clean
clean: ## Remove build output
	rm -rf bin coverage.out
