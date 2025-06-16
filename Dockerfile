FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o geeson-auth ./cmd/server

# Use a smaller image for the runtime
FROM alpine:latest

# Install CA certificates for HTTPS
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/geeson-auth .

# Set environment variables (these can be overridden at runtime)
ENV PORT=8087 \
    DB_HOST=mysql \
    DB_PORT=3306 \
    DB_USER=root \
    DB_PASS=password \
    DB_NAME=geeson_auth

# Expose the port the app runs on
EXPOSE 8087

# Run the binary
CMD ["./geeson-auth"]