FROM golang:1.20.5-alpine3.18

RUN apk add --no-cache \
    gcc \
    libc-dev \
    tzdata \
    ca-certificates

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

RUN go install github.com/cortesi/modd/cmd/modd@latest

CMD ["sh", "-c", "/go/bin/linux_amd64/modd -f modd.conf"]