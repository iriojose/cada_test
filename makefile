# Variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Name of the executable
BINARY_NAME=myapp

all: test build

build:
    $(GOBUILD) -o $(BINARY_NAME)

clean:
    $(GOCLEAN)
    rm -f $(BINARY_NAME)

test:
    $(GOTEST) -v ./...

run: build
    ./$(BINARY_NAME)

deps:
    $(GOGET) ./...