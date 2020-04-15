# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=app

all: test build
config/bindatafs/templates_bindatafs.go:
	$(GOBUILD) -o compile -v
	./compile compileStatic
	rm compile
build: config/bindatafs/templates_bindatafs.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v -tags 'bindatafs'
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)
