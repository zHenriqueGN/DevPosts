FROM golang:1.20-alpine

WORKDIR /build

COPY . .

RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o apiserver cmd/api/main.go
