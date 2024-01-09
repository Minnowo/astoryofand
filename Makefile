
@PHONY: run debug build_template

all: debug

build_template:
	templ generate

run:  | build_template
	@DEBUG=false LOG_LEVEL=1 go run cmd/main.go

debug: | build_template
	@DEBUG=true go run cmd/main.go
