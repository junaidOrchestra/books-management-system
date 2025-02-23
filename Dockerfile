# Use Go as the base image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project
COPY . .

# Build the application
RUN go build -o books-management-system ./cmd/main.go
CMD ["./books-management-system"]

## Use a lightweight image for production
#FROM alpine:latest
#
#WORKDIR /app
#
## Copy the built binary from the builder stage
#COPY --from=builder /app/books-management-system .
#
#RUN chmod +x /app/books-management-system
#
## Expose the port (change according to your app)
#EXPOSE 8080
#
## Run the application
#CMD ["ls", "/app/books-management-system"]
#CMD ["/app/books-management-system"]
