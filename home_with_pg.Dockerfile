
FROM golang:1.22.3-alpine AS builder

# Required stuff for building
RUN apk update && \
    apk add --no-cache --update gcc musl-dev make && \
    apk add --no-cache ca-certificates && \
    update-ca-certificates

WORKDIR /app

COPY . .

RUN make build_home_alpine_static_for_docker




# Yes, we are packaging postgres into our container
# This makes it super easy to have it as a db, no extra bs with containers and networking
FROM postgres:16.3-alpine

WORKDIR /app

ENV LOG_LEVEL=2
ENV DEBUG=false

COPY --from=builder /app/main .
COPY --from=builder /app/home_with_pg_entry.sh ./home-entrypoint.sh
RUN chmod +x ./home-entrypoint.sh

RUN mkdir conf && chown postgres:postgres conf && \
    mkdir enc  && chown postgres:postgres enc


# where the app writes encrypted data
VOLUME /app/enc
VOLUME /app/conf

EXPOSE 3000

USER postgres
CMD ["/app/home-entrypoint.sh"]


