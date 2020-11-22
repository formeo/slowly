FROM golang:1.13.8-alpine3.10 as build

ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"

WORKDIR /go/src/slowly
COPY . /go/src/slowly

RUN go mod download
RUN go mod vendor

RUN go build