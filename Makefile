
BIN_DIR  := bin

SITE_SRC := cmd/site/main.go
SITE_DST := $(BIN_DIR)/main

DECRYPT_SRC := cmd/decrypt/main.go
DECRYPT_DST := $(BIN_DIR)/decrypt

LDFLAGS ?= 

@PHONY: run debug build build_template build_save_docker bin_dir clean build_apline_static_for_docker setup_tailwind

all: debug

clean:
	rm -f $(SITE_DST) $(DECRYPT_DST)

bin_dir:
	mkdir -p $(BIN_DIR)

build_template:
	templ generate

build_css:
	./bin/tailwindcss-linux-x64 -i ./internal/assets/tailwind.css -o ./internal/assets/dist/css/tailwind.css --minify

setup_tailwind: | bin_dir
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.1/tailwindcss-linux-x64
	mv tailwindcss-linux-x64  $(BIN_DIR)
	chmod +x $(BIN_DIR)/tailwindcss-linux-x64

$(SITE_DST): $(SITE_SRC) | build_template bin_dir
	go build -ldflags "$(LDFLAGS)" -o $(SITE_DST) $(SITE_SRC) 

$(DECRYPT_DST): $(DECRYPT_SRC) | build_template bin_dir
	go build -ldflags "$(LDFLAGS)" -tags=include_private_key -o $(DECRYPT_DST) $(DECRYPT_SRC)

build: | $(SITE_DST) $(DECRYPT_DST) build_template bin_dir
	echo "Done"

build_apline_static_for_docker:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o main -ldflags "-s" $(SITE_SRC)

build_docker: $(SITE_SRC) | build_template bin_dir
	docker build -t astoryofand .

build_save_docker: $(SITE_SRC) | build_docker
	docker save -o bin/astoryofand.tar astoryofand:latest

run:  | build_template
	@DEBUG=false LOG_LEVEL=2 go run $(SITE_SRC)

debug: | build_template
	gofmt -w -s .
	goimports -w .
	@DEBUG=true go run $(SITE_SRC)

