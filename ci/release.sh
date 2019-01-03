#!/usr/bin/env bash

go mod vendor
packr -z
curl -sL https://git.io/goreleaser | bash
