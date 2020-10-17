
# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# Binary names
BINARY_NAME=valuez.so

all: build

race:
	$(GOBUILD) -race -trimpath -buildmode=plugin -o $(BINARY_NAME) -v ./...

build:
	$(GOBUILD) -trimpath -buildmode=plugin -o $(BINARY_NAME) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
