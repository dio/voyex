#!/usr/bin/env bash

go mod vendor
go get -u github.com/gobuffalo/packr/packr
git reset --hard
packr -z
curl -xL https://git.io/goreleaser | bash
