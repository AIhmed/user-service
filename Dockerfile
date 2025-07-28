# Build stage
FROM golang:1.24 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/server

# Runtime stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/user-service .
COPY --from=builder /app/config.yaml .
EXPOSE 8080 50051
CMD ["./user-service"]
