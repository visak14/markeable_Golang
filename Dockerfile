# Stage 1: Build the Go binary
FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules files (go.mod and go.sum) and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a smaller image with only the necessary runtime dependencies
FROM alpine:latest

# Install necessary CA certificates for HTTPS connections
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built Go binary from the build stage
COPY --from=builder /app/main .

# Copy the environment variable file (.env) if exists
COPY .env .env

# Expose the port the app will run on
EXPOSE 8000

# Run the Go binary
CMD ["./main"]
