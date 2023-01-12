FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest

COPY ./ /app/

CMD "air"