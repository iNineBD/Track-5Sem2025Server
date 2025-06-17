FROM golang:1.23.6 AS builder

WORKDIR /app
COPY . .
WORKDIR /app/src
RUN go mod tidy
RUN go build -o server main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/src/server ./server
EXPOSE 8080
CMD ["./server"]