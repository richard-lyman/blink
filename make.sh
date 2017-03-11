#!/bin/sh

set -e

GOPATH=`pwd` go build -o bin/blink -ldflags '-s' ./src/lithinos.com/blink/blink.go

