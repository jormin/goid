#!/bin/bash

export PATH=$PATH:/usr/local/go/bin
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct
go mod download
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goid main.go
