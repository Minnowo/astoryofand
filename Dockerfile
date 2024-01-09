FROM golang:latest AS builder

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags "-s" cmd/main.go




FROM alpine:latest

WORKDIR /app

# ENV LOG_LEVEL=4
ENV DEBUG=false

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static

EXPOSE 3000

CMD ["./main"]