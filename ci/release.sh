#!/usr/bin/env bash

go mod vendor
go get -u github.com/gobuffalo/packr/packr
packr -z
curl -sL https://git.io/goreleaser | bash
