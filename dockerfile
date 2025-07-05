FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod tidy

COPY .. ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

# Fase final
FROM alpine:3.20.3

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

EXPOSE 8181

ENTRYPOINT ["./main"] 