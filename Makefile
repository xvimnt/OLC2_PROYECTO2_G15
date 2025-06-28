# V Language Interpreter Makefile

.PHONY: all build clean test generate run run-arm coverage help build-arm test-v-flow install-qemu check-qemu

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
ARM_GCC=aarch64-linux-gnu-gcc
ARM_FLAGS=-static
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
	wsl -d Ubuntu $(ARM_GCC) $(ARM_FLAGS) -o $(ARM_OUTPUT) $(ARM_SOURCE)
 
# Install QEMU automatically using PowerShell script
install-qemu:
	@echo "=== Installing QEMU in WSL Ubuntu ==="
	@powershell -ExecutionPolicy Bypass -File install_qemu.ps1

# Check if QEMU is installed
check-qemu:
	@echo "Checking QEMU installation..."
	@powershell -Command " \
		try { \
			$$result = wsl -d Ubuntu bash -c 'command -v qemu-aarch64-static'; \
			if ($$result) { \
				Write-Host 'QEMU está instalado correctamente' -ForegroundColor Green; \
				wsl -d Ubuntu qemu-aarch64-static --version | Select-Object -First 1; \
			} else { \
				Write-Host 'QEMU no está instalado. Ejecuta: make install-qemu' -ForegroundColor Red; \
				exit 1; \
			} \
		} catch { \
			Write-Host 'Error verificando QEMU o WSL no disponible' -ForegroundColor Red; \
			exit 1; \
		}"

# Run the ARM executable using QEMU (with fallback)
run-arm:
	@echo "Running ARM executable with QEMU..."
	@powershell -Command " \
		$$wsl_path = ((Get-Location).Path.Replace('C:\', '/mnt/c/').Replace('D:\', '/mnt/d/').Replace('\', '/') + '/$(ARM_OUTPUT)'); \
		try { \
			wsl -d Ubuntu bash -c \"if command -v qemu-aarch64-static &> /dev/null; then /usr/bin/qemu-aarch64-static $$wsl_path; else echo 'ERROR: QEMU not found. Run: make install-qemu'; exit 1; fi\" \
		} catch { \
			Write-Host 'ERROR: WSL Ubuntu not available or QEMU not installed.'; \
			Write-Host 'Please run: make install-qemu'; \
			exit 1 \
		}"

# Run the full build, translate, build-arm, and run-arm flow for test.v
test-v-flow:
	@echo "--- [0/5] Checking QEMU installation ---"
	$(MAKE) check-qemu
	@echo "--- [1/5] Building the compiler ---"
	$(MAKE) build
	@echo "--- [2/5] Translating test.v to ARM assembly ---"
	powershell -Command "$$env:VLANG_DEBUG='true'; .\$(BINARY_NAME) translate test.v"
	@echo "--- [3/5] Building the ARM executable ---"
	$(MAKE) build-arm
	@echo "--- [4/5] Running the ARM executable with QEMU ---"
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
	@echo "  install-qemu - Install QEMU in WSL Ubuntu for ARM emulation"
	@echo "  check-qemu - Check if QEMU is installed and working"
	@echo "  run-arm    - Run the ARM executable using QEMU"
	@echo "  test-v-flow - Run the full flow for test.v (install-qemu, build, translate, build-arm, run-arm)"
	@echo "  example    - Run a specific example: make example basic (runs examples/basic.mylang)"
	@echo "  repl       - Start the REPL"
	@echo "  help       - Display this help information"

# Allow passing arguments to run target
%:
	@true
