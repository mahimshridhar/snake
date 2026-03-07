# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Binary name
BINARY_NAME=snake

# Main file
MAIN_FILE=main.go

# Default target
.PHONY: all
all: test build

# Build the game
.PHONY: build
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Run the game
.PHONY: run
run:
	$(GORUN) $(MAIN_FILE)

# Run with race detector (find bugs!)
.PHONY: race
race:
	$(GORUN) -race $(MAIN_FILE)

# Format code
.PHONY: fmt
fmt:
	$(GOFMT) ./...
	@echo "✅ Code formatted"

# Tidy dependencies
.PHONY: tidy
tidy:
	$(GOMOD) tidy
	@echo "✅ Dependencies tidied"

# Test (if you add tests later)
.PHONY: test
test:
	$(GOTEST) -v ./...

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	@echo "✅ Cleaned up"

# Install the game
.PHONY: install
install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Installed to $(GOPATH)/bin/$(BINARY_NAME)"

# Help menu
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build   - Build the binary"
	@echo "  make run     - Run the game"
	@echo "  make race    - Run with race detector"
	@echo "  make fmt     - Format code"
	@echo "  make tidy    - Tidy dependencies"
	@echo "  make test    - Run tests"
	@echo "  make clean   - Remove binary"
	@echo "  make install - Install to GOPATH/bin"
	@echo "  make help    - Show this help"