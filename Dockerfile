# Stage 1: build binary
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .

RUN go build -o main .

# Stage 2: runtime
FROM debian:bullseye-slim

WORKDIR /app
COPY --from=builder /app/main .

RUN chmod +x /app/main

ENV PORT=3000
EXPOSE 3000

CMD ["./main"]
