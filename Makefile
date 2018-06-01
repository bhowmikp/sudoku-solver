# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=sudoku-solver

all: build
build:
				$(GOBUILD) -o $(BINARY_NAME) -v
test:
				$(GOTEST) -v ./...
clean:
				rm -f $(BINARY_NAME)
run:
				./$(BINARY_NAME)
