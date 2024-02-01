# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=api-sekejap
BINARY_UNIX=$(BINARY_NAME)_unix

# Default target executed when no arguments are given to make.
default: build

# Builds the project
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Runs tests
test:
	$(GOTEST) -v ./...

# Cleans our project: deletes binaries
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Runs the project
run-http:
	$(GORUN) cmd/app/http/main.go
run-grpc:
	$(GORUN) cmd/app/grpc/main.go

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

.PHONY: default build test clean run-http run-grpc build-linux