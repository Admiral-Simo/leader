FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o webapp ./cmd/webapp/main.go

FROM debian:bullseye-slim

COPY --from=builder /app/webapp /usr/local/bin/webapp

EXPOSE 8080

CMD ["webapp"]
