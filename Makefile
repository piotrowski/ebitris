.PHONY: help build lint test

help: ## Show available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-10s %s\n", $$1, $$2}'

build: ## Build the binary
	go build ./cmd/ebitris

lint: ## Run linter
	golangci-lint run

test: ## Run all tests
	go test ./...
