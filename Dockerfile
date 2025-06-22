FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/validator main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/validator .

RUN chmod +x ./validator

ENTRYPOINT ["./validator"]
