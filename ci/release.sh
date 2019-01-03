#!/usr/bin/env bash

go mod vendor
go get github.com/gobuffalo/packr/packr
git reset --hard
packr -z
curl -sSL https://git.io/goreleaser | bash
