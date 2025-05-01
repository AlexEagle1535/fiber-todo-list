# Стадия 1 — сборка
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

# Стадия 2 — минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/main .
COPY .env .
COPY db/ ./db/  

RUN apk add --no-cache ca-certificates

EXPOSE 3000

CMD ["./main", "--migrate"]