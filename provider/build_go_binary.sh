#!/bin/bash

export GOPATH=$(pwd)

cd src/provider
go get
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s" -a -installsuffix cgo -o $GOPATH/provider

cd $GOPATH


