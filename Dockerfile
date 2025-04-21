# Etapa 1 - Build
FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server
RUN go build -o /main

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /main .

EXPOSE 8080
CMD ["./main"]
