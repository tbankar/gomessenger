#!/bin/bash

set -x
wget https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
mv dep-linux-amd64 /usr/bin/dep
chmod +x /usr/bin/dep
cd /go/src/gomessenger/apig/cmd
dep ensure -v
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apig .

