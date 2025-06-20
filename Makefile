# V Language Interpreter Makefile

.PHONY: all build clean test generate run coverage help

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=OLC2_PROYECTO1_G15
MAIN_PKG=./cmd/
GRAMMAR_DIR=./grammar
PARSER_DIR=./parser
TEST_DIR=./tests
COVERAGE_FILE=coverage.out
ANTLR_JAR=antlr-4.12.0-complete.jar
ANTLR_URL=https://www.antlr.org/download/antlr-4.12.0-complete.jar

# Default target
all: download-antlr generate build

# Build the application
build:
	@echo "Building application..."
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PKG)

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f $(COVERAGE_FILE)
	rm -rf $(PARSER_DIR)/*

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run test coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -coverprofile=$(COVERAGE_FILE) ./...
	$(GOCMD) tool cover -html=$(COVERAGE_FILE)

# Download ANTLR if not already present
download-antlr:
	@echo "Checking for ANTLR jar..."
	@if [ ! -f $(ANTLR_JAR) ]; then \
		echo "Downloading ANTLR jar..."; \
		curl -O $(ANTLR_URL); \
	else \
		echo "ANTLR jar already exists."; \
	fi

# Generate parser from grammar
generate:
	@echo "Generating parser from grammar..."
	@mkdir -p $(PARSER_DIR)
	@echo "Running ANTLR to generate Go code..."
	java -Xmx500M -cp $(ANTLR_JAR) org.antlr.v4.Tool -Dlanguage=Go -visitor -o $(PARSER_DIR) $(GRAMMAR_DIR)/VLangCherry.g4
	@if [ -d $(PARSER_DIR)/grammar ]; then \
	    mv $(PARSER_DIR)/grammar/* $(PARSER_DIR)/; \
	    rmdir $(PARSER_DIR)/grammar; \
	fi

# Fix imports in generated files
fix-imports:
	@echo "Fixing imports in generated files..."
	$(GOCMD) fmt $(PARSER_DIR)/...

# Update dependencies
deps:
	@echo "Updating dependencies..."
	$(GOMOD) tidy

# Run the application
run:
	@echo "Running application..."
	./$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

# Run a specific example
example:
	@echo "Running example: $(filter-out $@,$(MAKECMDGOALS))"
	./$(BINARY_NAME) run ./examples/$(filter-out $@,$(MAKECMDGOALS)).mylang

# Run the REPL
repl:
	@echo "Starting REPL..."
	./$(BINARY_NAME) repl

# Display help information
help:
	@echo "V Language Interpreter - Makefile targets:"
	@echo "  all        - Download ANTLR, generate parser and build the application"
	@echo "  build      - Build the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  coverage   - Run tests with coverage report"
	@echo "  generate   - Generate parser from grammar"
	@echo "  deps       - Update Go dependencies"
	@echo "  run        - Run the application"
	@echo "  example    - Run a specific example: make example basic (runs examples/basic.mylang)"
	@echo "  repl       - Start the REPL"
	@echo "  help       - Display this help information"

# Allow passing arguments to run target
%:
	@true
