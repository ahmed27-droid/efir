FROM golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git make

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o broadcast-service ./cmd/app

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/broadcast-service .

EXPOSE 8082

CMD ["./broadcast-service"]
