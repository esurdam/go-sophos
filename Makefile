# Go parameters
GOCMD=go
GOFMT=gofmt
GOIMPORTS=goimports
GOTEST=$(GOCMD) test
GENOUTPUT="types/generated.go"

all: build fmt test
build:
	([[ -f $GENOUTPUT ]] && rm $GENOUTPUT || echo )
	$(GOCMD) run bin/gen.go
fmt:
	$(GOFMT) -s -w .
	$(GOIMPORTS) -w .
test:
	$(GOTEST) -race -v -coverprofile=coverage.txt -covermode=atomic
clean:
	$(GOCLEAN)