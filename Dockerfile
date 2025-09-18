# Build stage
FROM golang:1.24.6 AS builder

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build binary from cmd/api/main.go
RUN go build -o main ./cmd/api

# Runtime stage (small, secure image)
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/main .

# Expose port for clarity (optional, Compose handles it anyway)
EXPOSE 8080

# Run the Go app
CMD ["./main"]
