#!/usr/bin/env bash

go get -u github.com/gobuffalo/packr/packr
packr -z
go mod download