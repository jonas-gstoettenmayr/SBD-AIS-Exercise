#!/bin/sh
"$(go env GOPATH)/bin/swag" init --pd --st -g main.go --dir ./ -o docs
