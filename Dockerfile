FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o webapp ./cmd/webapp/main.go

FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y \
    libc6 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/webapp /usr/local/bin/webapp
COPY --from=builder /app/public /app/public

EXPOSE 8080

CMD ["webapp"]
