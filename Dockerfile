FROM golang:1.17.1-alpine as builder

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN apk update && apk add git --no-cache ca-certificates apache2-utils
RUN go get github.com/boazsoftwares/Boaz.Api.Go/cmd/boazApi

RUN cd /build && git clone https://github.com/boazsoftwares/Boaz.Api.Go.git
RUN cd /build/Boaz.Api.Go/cmd/boazApi && go build main.go

EXPOSE 8080

ENTRYPOINT [ "boazApi" ]