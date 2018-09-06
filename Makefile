# Go parameters
GOCMD=go
GOFMT=gofmt
GOIMPORTS=goimports
GOTEST=$(GOCMD) test
GENOUTPUT="types/generated.go"

all: build test
build:
	([[ -f $GENOUTPUT ]] && rm $GENOUTPUT || echo )
	$(GOCMD) run bin/gen.go
	$(GOFMT) -s -w .
	$(GOIMPORTS) -w .
test:
	$(GOTEST) -v .
clean:
	$(GOCLEAN)