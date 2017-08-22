GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./... | grep -v /vendor/)
BUILDDIR = build/

default: build

build: $(GOFILES)
	mkdir -p ${BUILDDIR}
	go build -o ${BUILDDIR}/http .

.PHONY: build-docker
build-docker:
	docker build -t ${DOCKER_TAG} .

.PHONY: clean
clean:
	rm -rf ${BUILDDIR}
