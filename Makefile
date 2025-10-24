# Gobrou Static Blog Generator Makefile

# Variables
BINARY_NAME=gobrou
BUILD_DIR=.
TEST_DIR=testing/painhardcore.ru
DIST_DIR=dist

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Build the application
.PHONY: build
build: ## Build the Gobrou binary
	@echo "Building $(BINARY_NAME)..."
	@go generate ./...
	@go build -o $(BINARY_NAME) $(BUILD_DIR)
	@echo "Build complete: $(BINARY_NAME)"

# Clean build artifacts
.PHONY: clean
clean: ## Clean build artifacts and generated files
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(DIST_DIR)
	@rm -f templates/*.qtpl.go
	@echo "Clean complete"

# Build and run for testing
.PHONY: test-build
test-build: ## Build the application and generate static site for testing
	@echo "Building for testing in $(TEST_DIR)..."
	@cd $(TEST_DIR) && \
		go generate ./... && \
		go build -o ../../$(BINARY_NAME) ../../ && \
		./../../$(BINARY_NAME) && \
		echo "Test build complete in $(TEST_DIR)/$(DIST_DIR)"

# Run development server for testing
.PHONY: test-run
test-run: ## Run development server for testing (requires test-build first)
	@echo "Starting development server for testing..."
	@cd $(TEST_DIR) && \
		if [ ! -f "../../$(BINARY_NAME)" ]; then \
			echo "Binary not found. Run 'make test-build' first."; \
			exit 1; \
		fi && \
		./../../$(BINARY_NAME) -dev

# Build and run for testing in one command
.PHONY: test
test: test-build test-run ## Build and run development server for testing

# Install dependencies
.PHONY: deps
deps: ## Install Go dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies installed"

# Generate templates
.PHONY: generate
generate: ## Generate template files
	@echo "Generating templates..."
	@go generate ./...
	@echo "Template generation complete"

# Run tests
.PHONY: test-unit
test-unit: ## Run unit tests
	@echo "Running unit tests..."
	@go test ./...

# Format code
.PHONY: fmt
fmt: ## Format Go code
	@echo "Formatting code..."
	@go fmt ./...

# Lint code
.PHONY: lint
lint: ## Lint Go code
	@echo "Linting code..."
	@go vet ./...

# All-in-one development setup
.PHONY: dev-setup
dev-setup: deps generate build ## Complete development setup

# Production build
.PHONY: prod-build
prod-build: clean deps generate build ## Production build with clean dependencies
