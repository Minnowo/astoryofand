FROM golang:latest

WORKDIR /app

COPY . .

# ENV LOG_LEVEL=4
ENV DEBUG=false

RUN go build -o main -ldflags "-s" cmd/main.go

# remove the source code
RUN find . -type f -name "*.go" -delete
RUN find . -type f -name "*.templ" -delete
RUN find . -type d -empty -delete

EXPOSE 3000

CMD ["./main"]