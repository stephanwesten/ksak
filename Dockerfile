# syntax=docker/dockerfile:1

#FROM golang:1.19-alpine
FROM golang:1.19

RUN mkdir -p /app

WORKDIR /app

COPY bin/ksak .

#CMD ./ksak
ENTRYPOINT /app/ksak