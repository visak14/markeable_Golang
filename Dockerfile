# Stage 1 - Builder
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .  # This copies all files, including .env, into the build context
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2 - Final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY .env .env  # This copies the .env file into the container
CMD ["./main"]
