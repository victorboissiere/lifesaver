FROM golang:1.10.2-alpine3.7

RUN apk --update add git && \
    go get gopkg.in/yaml.v2

COPY ./lifesaver.go ./
COPY ./installer /go/installer/

RUN go build lifesaver.go

FROM ubuntu:bionic

RUN apt update && apt install wget sudo git -y
COPY --from=0 /go/lifesaver /tmp/lifesaver
RUN chmod +x /tmp/lifesaver

USER root


