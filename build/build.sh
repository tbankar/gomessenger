#!/bin/bash

gateway=$1
server=$2

wget https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64
mv dep-linux-amd64 /usr/bin/dep
chmod +x /usr/bin/dep
cd /go/src/gomessenger
dep init
dep ensure -v

if [ $gateway -eq 1 ]; then
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apigw ./apig/cmd
elif [ $server -eq 1 ]; then
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o msngr ./server/cmd/msngr
fi

