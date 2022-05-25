#!/usr/bin/env bash
# linux server build imagersync-service images
set -x
set -e

imageTagVersion="v1.0"
registrieAddress="harbor.xxx.com/devops"
servicename=imagersync_service
pkgname=imagersync_service_bin

export GOARCH=amd64
export GOOS=linux
export GCCGO=gc

go build -o build/${pkgname} main.go
chmod u+x build/${pkgname}

docker build -f build/dockerfile -t ${registrieAddress}/${servicename}:${imageTagVersion} .
docker push ${registrieAddress}/${servicename}:${imageTagVersion}
