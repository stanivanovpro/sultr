FROM golang:1.19-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

ENV GIN_MODE=release

COPY . .
RUN go mod download
RUN go build -o app main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/app .

EXPOSE 8080
ENTRYPOINT ["./app"]