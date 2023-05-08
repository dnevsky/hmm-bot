FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN chmod +x wait-for-it.sh

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o hmm-bot cmd/main.go
 
FROM alpine

RUN apk update && apk add --no-cache bash

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/hmm-bot /app/hmm-bot
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/wait-for-it.sh /app/wait-for-it.sh

WORKDIR /app

CMD ["./hmm-bot"]