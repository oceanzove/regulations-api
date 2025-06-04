FROM --platform=linux/amd64 golang:1.23-bullseye AS builder

WORKDIR /app
COPY . .
COPY configs/config.json /app/configs/config.json

RUN go mod download
RUN go build -o main ./cmd/main.go

FROM debian:bullseye

WORKDIR /app

RUN apt-get update && apt-get install -y netcat-openbsd && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/main .
COPY .env .env
COPY --from=builder /app/configs/config.json ./configs/config.json
COPY --from=builder /app/migrations ./migrations

EXPOSE 25504
CMD ["./main"]
