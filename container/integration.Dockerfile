FROM golang:1.18

WORKDIR /app

COPY . .

ENTRYPOINT [ "/app/tests/entrypoint_integration.sh" ]