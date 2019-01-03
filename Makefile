GO ?= go
PACKR := $(or ${shell which packr},packr)
COMMIT = $(shell git rev-parse --short HEAD)
VERSION = $(shell git describe --tags)
LDFLAGS = "-s -w -X main.commit=$(COMMIT) -X main.version=$(VERSION)"

all: build

build: $(PACKR) main.go
	@$(PACKR) -z
	@$(GO) build -ldflags $(LDFLAGS) -o bin/voyex main.go

$(PACKR):
	@go get github.com/gobuffalo/packr/packr

clean:
	rm -fr bin

.PHONY: build
