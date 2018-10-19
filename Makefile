# Go parameters
GOCMD=go

all: build
build:
				$(GOCMD)fmt -w .
				$(GOCMD) build
lint:
				$(GOCMD)fmt -w .
test:
				$(GOCMD) test
clean:
				$(GOCMD) clean
run:
				$(GOCMD) run app
