FROM golang:1.13-alpine

RUN mkdir /go/src/app
ADD . /go/src/app/
WORKDIR /go/src/app
RUN  go mod download && \
    go build -o server .