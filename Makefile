# Go parameters
GOCMD=go
GOFMT=gofmt
GOIMPORTS=goimports
GOTEST=$(GOCMD) test

all: build test
build:
	([[ -f generated.go ]] && rm generated.go)
	$(GOCMD) run bin/gen.go
	$(GOFMT) -s -w .
	$(GOIMPORTS) -w .
test:
	$(GOTEST) .
clean:
	$(GOCLEAN)