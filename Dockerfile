FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod ./
# Copy all project files
COPY . .

# Build the application
RUN go build -o ascii-studio-server ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app

# Copy the binary
COPY --from=builder /app/ascii-studio-server .
# Copy static files and banners
COPY --from=builder /app/web ./web
COPY --from=builder /app/banners ./banners

# Create data directory for JSON storage
RUN mkdir -p data

EXPOSE 8087

CMD ["./ascii-studio-server"]
