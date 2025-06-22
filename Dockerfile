FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o /app/bin/validator main.go

FROM alpine:latest

# Cria usuário não-root
RUN adduser -D appuser

WORKDIR /home/appuser

COPY --from=builder /app/bin/validator /usr/local/bin/validator

USER appuser

HEALTHCHECK --interval=30s --timeout=3s CMD ["/usr/local/bin/validator", "--cpf", "11111111111"]

ENTRYPOINT ["/usr/local/bin/validator"]
