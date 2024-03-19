FROM golang:1.22.1-alpine AS builder

# Required stuff for building
RUN apk update && \
    apk add --no-cache --update gcc musl-dev make && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

# Create a user so that the image doens't run as root
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "100001" \
    "appuser"

WORKDIR /app

COPY . .

RUN make build_site_alpine_static_for_docker




FROM alpine:latest

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

# ENV LOG_LEVEL=4
ENV DEBUG=false

COPY --from=builder /app/main .

RUN mkdir conf && chown appuser:appuser conf && \
    mkdir enc  && chown appuser:appuser enc

USER appuser:appuser

# where the app writes encrypted data
VOLUME /app/enc
VOLUME /app/conf

EXPOSE 3000

CMD ["./main"]
