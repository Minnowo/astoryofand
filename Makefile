
@PHONY: run debug build build_template build_save_docker bin_dir

all: debug

bin_dir:
	mkdir -p bin

build_template:
	templ generate

build: cmd/main.go | build_template bin_dir
	go build -o bin/main cmd/main.go
	cp -r static bin/static

build_docker: cmd/main.go | build_template bin_dir
	docker build -t astoryofand .

build_save_docker: cmd/main.go | build_docker
	docker save -o bin/astoryofand.tar astoryofand:latest

run:  | build_template
	@DEBUG=false LOG_LEVEL=2 go run cmd/main.go

debug: | build_template
	@DEBUG=true go run cmd/main.go
