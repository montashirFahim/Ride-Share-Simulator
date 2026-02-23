FROM golang:1.25.7-alpine AS builder
WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY .env /app
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o app

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app", "server"]