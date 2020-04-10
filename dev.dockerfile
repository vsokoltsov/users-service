FROM golang:1.13

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go get github.com/cespare/reflex && \ 
    go get -u github.com/pressly/goose/cmd/goose && \
    go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app .
