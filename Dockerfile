FROM golang:1.15

RUN apt-get update && apt-get install -y git curl protobuf-compiler

WORKDIR /go/src

ENV GO111MODULE on
RUN go get -u github.com/golang/protobuf/protoc-gen-go
ENV PATH $PATH:$(go env GOPATH)/bin
