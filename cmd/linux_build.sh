#!/usr/bin/env bash
# linux server build go-default-service images
set -x
set -e

imageTagVersion="v0.1"
registrieAddress="harbor.xxx.com"

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -o build/go-default-service go_default.go
chmod u+x build/go-default-service

docker build -f build/dockerfile -t ${registrieAddress}/devops/go-default-service:${imageTagVersion} .
docker push ${registrieAddress}/devops/go-default-service:${imageTagVersion}
