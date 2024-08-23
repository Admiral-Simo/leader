FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o webapp ./cmd/webapp/main.go

FROM alpine:latest

RUN apk add --no-cache libc6-compat

COPY --from=builder /app/webapp /usr/local/bin/webapp

EXPOSE 8080

CMD ["webapp"]
