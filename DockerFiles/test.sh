#!/bin/bash
echo "Start: testing the program"
mkdir -p /etc/api/
cp /go/src/github.com/rebelit/rpIoT/api_config.json /etc/api/api_config.json
cd /go/src/github.com/rebelit/rpIoT
go test
echo "Done: building the program"
