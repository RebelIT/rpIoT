FROM golang:1.12beta2-stretch

RUN mkdir -p /go/src/github.com/rebelit/rpIoT

ADD DockerFiles/run.sh /run.sh
RUN chmod +x /run.sh

ADD DockerFiles/test.sh /test.sh
RUN chmod +x /test.sh

ADD DockerFiles/build.sh /build.sh
RUN chmod +x /build.sh