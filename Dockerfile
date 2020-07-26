# Install necessary binaries
FROM golang:1.13-alpine as binaries

RUN apk add git gcc g++
#  make unzip gettext rsync gcc g++
RUN go get github.com/cespare/reflex && \ 
    go get google.golang.org/grpc && \
    go get github.com/golang/protobuf/protoc-gen-go && \
#     go get github.com/golang/protobuf && \
    go get -u github.com/pressly/goose/cmd/goose

# Base app image
FROM golang:1.13-alpine

RUN apk add protobuf gcc g++
COPY --from=binaries /go/bin /go/bin
WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app .