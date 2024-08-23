# Start from the official Golang base image
FROM golang:1.22.4 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifest
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o webapp ./cmd/webapp/main.go

# Start a new stage from scratch
FROM debian:bullseye-slim

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/webapp /usr/local/bin/webapp

# Expose port 80 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["webapp"]
