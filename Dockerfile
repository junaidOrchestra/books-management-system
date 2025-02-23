# Use the official Golang image as a builder
FROM golang:1.23 AS builder

WORKDIR /app

# Copy everything and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN go build -o books-management-system ./cmd/main.go

# Use a minimal base image for running the app
FROM alpine:latest

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/books-management-system .

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./books-management-system"]
