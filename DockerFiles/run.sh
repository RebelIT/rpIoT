#!/bin/bash
echo "Start: running the program"
mkdir -p /etc/api/ssl
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -subj "/C=US/ST=Minnesota/L=MyCity/O=pi/CN=www.MyCoolAwesomtSite.com" -keyout /etc/api/key.key -out /etc/api/cert.cert
cp /go/src/github.com/rebelit/rpIoT/api_config.json /etc/api/api_config.json
cd /go/src/github.com/rebelit/rpIoT
go build -o main .
./main
