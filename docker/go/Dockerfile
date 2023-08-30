FROM golang:1.20.0-alpine3.16

ENV CGO_ENABLED=1

RUN apk update
RUN apk upgrade
RUN apk add --update figlet
RUN apk add build-base tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    echo "Asia/Tokyo" > /etc/timezone

# Go install Pack
RUN go install github.com/cespare/reflex@latest
