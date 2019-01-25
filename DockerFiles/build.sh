#!/bin/bash
echo "Start: building the program"
cd /go/src/github.com/rebelit/rpIoT
go build -o main .
echo "Done: building the program"