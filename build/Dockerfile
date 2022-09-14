FROM golang:1.18.5-alpine3.15 AS base
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on GOPROXY=https://goproxy.cn
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download
COPY . $WORKDIR
RUN  go build -o /usr/local/bin/gosky  main.go && go clean -cache

FROM alpine:3.15
COPY --from=base /usr/local/bin/gosky /usr/local/bin/
COPY local.config.yaml /
COPY production.config.yaml /