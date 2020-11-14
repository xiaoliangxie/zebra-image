#!/usr/bin/env bash
#mac环境下编译linux环境下的运行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go