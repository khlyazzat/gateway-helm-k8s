FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY services/api-gateway ./services/api-gateway
COPY utils ./utils
COPY pkg ./pkg
COPY values ./values

WORKDIR /app/services/api-gateway

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/services/api-gateway/app .

ENTRYPOINT ["./app"]
# ENTRYPOINT ["sh", "-c", "env && ./app"]

