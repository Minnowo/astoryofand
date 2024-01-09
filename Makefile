
@PHONY: run debug build build_template

all: debug

build_template:
	templ generate

build: cmd/main.go | build_template
	go build -o main cmd/main.go

run:  | build_template
	@DEBUG=false LOG_LEVEL=1 go run cmd/main.go

debug: | build_template
	@DEBUG=true go run cmd/main.go
