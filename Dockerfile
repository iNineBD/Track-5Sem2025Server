FROM golang:1.24 AS builder

WORKDIR /app
COPY . .
WORKDIR /app/src
RUN go mod tidy
RUN go build -o server main.go

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/src/server ./server
EXPOSE 8080
CMD ["./server"]