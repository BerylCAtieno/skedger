GO=go
SERVER_CMD=cmd/server/main.go

# Default target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run        - Start the server"
	@echo "  make dev        - Start the server in development mode with air"
	@echo "  make test       - Run all tests"

# Start the server
.PHONY: run
run:
	$(GO) run $(SERVER_CMD)

# Start development server
.PHONY: dev
dev:
	air

# Run tests
.PHONY: test
test:
	$(GO) test ./...
