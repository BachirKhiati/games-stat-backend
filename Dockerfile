FROM golang:1.17-alpine

RUN apk update && apk add git bash build-base curl

RUN mkdir -p /app
WORKDIR /app
ADD . /app

ENV GO111MODULE=on
#development only
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

EXPOSE 8000