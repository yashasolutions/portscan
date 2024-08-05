# Define the binary name
BINARY_NAME = port_scanner

# Define the source file
SOURCE = port_scanner.go

# Default target to build the binary
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) $(SOURCE)

# Run the program with default arguments
run: build
	./$(BINARY_NAME) -host=example.com -ports=80,443

# Clean the build artifacts
clean:
	rm -f $(BINARY_NAME)

# Phony targets to prevent conflicts with file names
.PHONY: all build run clean

