FROM golang:1.18-alpine

RUN apk update && apk add --no-cache musl-dev gcc git build-base

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]