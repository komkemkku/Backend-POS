# Build stage
FROM golang:1.23-alpine AS builder

# Install git and ca-certificates (for any https calls)
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Final stage
FROM alpine:latest

# Install ca-certificates for https calls
RUN apk --no-cache add ca-certificates tzdata

# Set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy any necessary files (like sample_data.sql if needed)
COPY --from=builder /app/sample_data.sql ./

# Expose port (Railway will override this with $PORT)
EXPOSE 8080

# Command to run
CMD ["./main"]
