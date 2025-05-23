# Dockerfile
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o load-tester .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/load-tester .

CMD ["./load-tester"]
