FROM golang:1.19.0-alpine3.16 AS builder

WORKDIR /app
COPY . .

RUN cat ./cmd/server/main.go


RUN CGO_ENABLED=0 GOOS=linux go build -o build/server ./cmd/server/main.go

# Create a data folder if it doesn't exist
RUN mkdir data; exit 0

FROM scratch AS production

WORKDIR /server

COPY --from=builder /app/prod.config.toml config.toml
# Copy data so, that it keeps the same path as in the dev container
COPY --from=builder /app/data data
COPY --from=builder /app/build/server .

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./server"]
