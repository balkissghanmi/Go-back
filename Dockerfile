# ===== Stage 1: Builder =====
FROM golang:1.25 as builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o go-crud-app .

# ===== Stage 2: Runtime =====
FROM golang:1.25

WORKDIR /app

# Copy binary
COPY --from=builder /app/go-crud-app .

# Expose port
EXPOSE 8080

CMD ["./go-crud-app"]
