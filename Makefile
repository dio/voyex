GO ?= go
COMMIT = $(shell git rev-parse --short HEAD)
VERSION = $(shell git describe --tags)
LDFLAGS = "-s -w -X main.commit=$(COMMIT) -X main.version=$(VERSION)"

all: build

build: main.go
	@$(GO) build -ldflags $(LDFLAGS) -o bin/voyex main.go

clean:
	rm -fr bin

.PHONY: build
