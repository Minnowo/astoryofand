FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN make build_apline_static_for_docker




FROM alpine:latest

WORKDIR /app

# ENV LOG_LEVEL=4
ENV DEBUG=false

COPY --from=builder /app/main .

# where the app writes encrypted data
VOLUME /app/enc

EXPOSE 3000

CMD ["./main"]
