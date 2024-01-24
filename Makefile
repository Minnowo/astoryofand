
BIN_DIR  := bin

SITE_SRC := cmd/site/main.go
SITE_DST := $(BIN_DIR)/main

DECRYPT_SRC := cmd/decrypt/main.go
DECRYPT_DST := $(BIN_DIR)/decrypt


@PHONY: run debug build build_template build_save_docker bin_dir

all: debug

bin_dir:
	mkdir -p $(BIN_DIR)

build_template:
	templ generate

$(SITE_DST): $(SITE_SRC) | build_template bin_dir
	go build -o $(SITE_DST) $(SITE_SRC)

$(DECRYPT_DST): $(DECRYPT_SRC) | build_template bin_dir
	go build -tags=include_private_key -o $(DECRYPT_DST) $(DECRYPT_SRC)

build: | $(SITE_DST) $(DECRYPT_DST) build_template bin_dir
	cp -r static bin/static

build_docker: cmd/main.go | build_template bin_dir
	docker build -t astoryofand .

build_save_docker: cmd/main.go | build_docker
	docker save -o bin/astoryofand.tar astoryofand:latest

run:  | build_template
	@DEBUG=false LOG_LEVEL=2 go run $(SITE_SRC)

debug: | build_template
	@DEBUG=true go run $(SITE_SRC)
