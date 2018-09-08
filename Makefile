# Go parameters
GOCMD=go
GOFMT=gofmt
GOIMPORTS=goimports
GOTEST=$(GOCMD) test
GENOUTPUT="types/generated.go"

all: gen test
gen: build fmt
build:
	(rm types/*.go || echo "")
	(rm nodes/*.go || echo "")
	$(GOCMD) run bin/gen.go
fmt:
	$(GOFMT) -s -w .
	$(GOIMPORTS) -w .
test:
	$(GOTEST) -race -v -coverprofile=coverage.txt -covermode=atomic
clean:
	$(GOCLEAN)