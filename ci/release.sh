#!/usr/bin/env bash

go get -u github.com/gobuffalo/packr/packr
packr -z
go mod vendor
curl -sL https://git.io/goreleaser | bash
