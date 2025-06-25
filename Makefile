# V Language Interpreter Makefile

.PHONY: all build clean test generate run run-arm coverage help build-arm test-v-flow

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=OLC2_PROYECTO2_G15.exe
MAIN_PKG=./cmd/
GRAMMAR_DIR=./grammar
PARSER_DIR=./parser
TEST_DIR=./tests
COVERAGE_FILE=coverage.out
ANTLR_JAR=antlr-4.12.0-complete.jar
ANTLR_URL=https://www.antlr.org/download/antlr-4.12.0-complete.jar

# ARM Cross-compilation parameters
ARM_GCC=aarch64-none-linux-gnu-gcc
ARM_FLAGS=-nostartfiles
ARM_OUTPUT=output.exe
ARM_SOURCE=output.s

# Default target
all: download-antlr generate build

# Build the application
build:
	@echo "Building application..."
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PKG)

# Build the ARM executable from assembly
build-arm:
	@echo "Building ARM executable from $(ARM_SOURCE)..."
	$(ARM_GCC) $(ARM_FLAGS) -o $(ARM_OUTPUT) $(ARM_SOURCE)

# Run the ARM executable using QEMU
run-arm:
	@echo "Running ARM executable with QEMU..."
	wsl /usr/bin/qemu-aarch64-static -L "/mnt/c/Program Files (x86)/Arm GNU Toolchain aarch64-none-linux-gnu/14.2 rel1/aarch64-none-linux-gnu/libc" "/mnt/c/Users/matri/OneDrive/Documents/USAC/OLC2_PROYECTO2_G15/output.exe"

# Run the full build, translate, build-arm, and run-arm flow for test.v
test-v-flow:
	@echo "--- [1/4] Building the compiler ---"
	$(MAKE) build
	@echo "--- [2/4] Translating test.v to ARM assembly ---"
	.\$(BINARY_NAME) translate test.v
	@echo "--- [3/4] Building the ARM executable ---"
	$(MAKE) build-arm
	@echo "--- [4/4] Running the ARM executable with QEMU ---"
	$(MAKE) run-arm

# Clean build artifacts
clean:
	@echo "Cleaning..."
	if exist $(BINARY_NAME) del $(BINARY_NAME)
	if exist $(ARM_OUTPUT) del $(ARM_OUTPUT)
	if exist $(COVERAGE_FILE) del $(COVERAGE_FILE)
	if exist $(PARSER_DIR) rmdir /s /q $(PARSER_DIR)

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
	@powershell -Command "if (-not (Test-Path $(ANTLR_JAR))) { Write-Host 'Downloading ANTLR jar...'; Invoke-WebRequest -Uri $(ANTLR_URL) -OutFile $(ANTLR_JAR) } else { Write-Host 'ANTLR jar already exists.' }"

# Generate parser from grammar
generate:
	@echo "Generating parser from grammar..."
	@if not exist "$(PARSER_DIR)" mkdir "$(PARSER_DIR)"
	@echo "Running ANTLR to generate Go code..."
	java -Xmx500M -cp $(ANTLR_JAR) org.antlr.v4.Tool -Dlanguage=Go -visitor -o $(PARSER_DIR) $(GRAMMAR_DIR)/VLangCherry.g4

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
	.\$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

# Run a specific example
example:
	@echo "Running example: $(filter-out $@,$(MAKECMDGOALS))"
	.\$(BINARY_NAME) run .\examples\$(filter-out $@,$(MAKECMDGOALS)).mylang

# Run the REPL
repl:
	@echo "Starting REPL..."
	.\$(BINARY_NAME) repl

# Display help information
help:
	@echo "V Language Interpreter - Makefile targets:"
	@echo "  all        - Download ANTLR, generate parser and build the application"
	@echo "  build      - Build the application"
	@echo "  build-arm  - Build the ARM executable from assembly"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  coverage   - Run tests with coverage report"
	@echo "  generate   - Generate parser from grammar"
	@echo "  deps       - Update Go dependencies"
	@echo "  run        - Run the application"
	@echo "  run-arm    - Run the ARM executable using QEMU"
	@echo "  test-v-flow - Run the full flow for test.v (build, translate, build-arm, run-arm)"
	@echo "  example    - Run a specific example: make example basic (runs examples/basic.mylang)"
	@echo "  repl       - Start the REPL"
	@echo "  help       - Display this help information"

# Allow passing arguments to run target
%:
	@true
