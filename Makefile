# Makefile for sheet CLI
GO = go
BINARY = sheet
SOURCES = $(wildcard cmd/sheet/*.go)

# Build for all platforms
all: windows linux

# Build for Windows
windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o bin/$(BINARY)-windows-amd64.exe $(SOURCES)

# Build for Linux
linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o bin/$(BINARY)-linux-amd64 $(SOURCES)

# Clean build artifacts
clean:
	rm -rf bin/

# Create bin directory
bin:
	mkdir -p bin