FROM golang:1.13

RUN mkdir /go/src/app
ADD . /go/src/app/
WORKDIR /go/src/app
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go build -o server .