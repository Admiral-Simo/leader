FROM golang:1.22.4 AS builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GIN_MODE=release

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o webapp -ldflags="-w -s" ./cmd/webapp/main.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/webapp /usr/local/bin/webapp

COPY ./public /app/public

COPY .env /app/.env

EXPOSE 8080

CMD ["webapp"]
