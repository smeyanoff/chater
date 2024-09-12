# Dockerfile
FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./backend main.go

# Используем минимальный образ для финального контейнера
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/backend .
COPY --from=builder /app/internal/config/cors-config.yaml ./config
COPY --from=builder /app/docs ./docs

CMD ["./backend"]