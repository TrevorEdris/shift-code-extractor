FROM golang:1.18-alpine AS builder

RUN apk update && apk add musl-dev gcc build-base ca-certificates

WORKDIR /app

ARG SERVICE

COPY . .
RUN go build -ldflags "-linkmode external -extldflags \"-static\" -s -w $LDFLAGS" -o the-binary cmd/${SERVICE}/main.go

# Copy the binary from "builder" into a scratch container to reduce the overall size of the image
FROM scratch AS final

ENTRYPOINT ["/app/the-binary"]
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app/the-binary /app/the-binary
