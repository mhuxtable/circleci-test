GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./... | grep -v /vendor/)

default: build

build: $(GOFILES)
	mkdir -p build/
	go build -o build/http .
