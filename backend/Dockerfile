# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies for CGO (required by go-sqlite3)
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with CGO enabled for SQLite (dynamic linking)
RUN CGO_ENABLED=1 go build -o api ./cmd/api

# Runtime stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata sqlite-libs

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/api .

# Copy migrations
COPY migrations ./migrations

# Create data directory for SQLite
RUN mkdir -p /data

# Set default environment variables
ENV PORT=8080
ENV DATABASE_PATH=/data/voca.db
ENV ENVIRONMENT=production

EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

CMD ["./api"]
