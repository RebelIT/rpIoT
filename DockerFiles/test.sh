#!/bin/bash
echo "Start: testing the program"
mkdir -p /etc/api/ssl
openssl req -new -newkey rsa:4096 -days 365 -nodes -x509 -subj "/C=US/ST=Minnesota/L=MyCity/O=pi/CN=www.MyCoolAwesomtSite.com" -keyout /etc/api/ssl/key.key -out /etc/api/ssl/cert.cert
cp /go/src/github.com/rebelit/rpIoT/api_config.json /etc/api/api_config.json
cd /go/src/github.com/rebelit/rpIoT
go test
echo "Done: building the program"
