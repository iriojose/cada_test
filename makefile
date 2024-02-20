# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the executable
BINARY_NAME=api

all: test build

build:
	$(GOBUILD) -o ./bin/ ./cmd/$(BINARY_NAME)

clean:	
	$(GOCLEAN)
		rm -f ./bin/$(BINARY_NAME)

test:
	$(GOTEST) -v ./...

run:
	./bin/$(BINARY_NAME)

deps:
	$(GOGET) ./...