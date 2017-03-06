#!/usr/bin/env bash

echo "Installing go dep..."
go get -u github.com/golang/dep/...

echo "Installing go deps..."
dep ensure

echo "Building mcrouter-kuberentes-provisioner..."
go build -a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo -o mcrouter-kuberentes-provisioner .
