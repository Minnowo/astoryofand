
VERSION := ${shell cat ./internal/assets/version.txt}

BIN_DIR  := bin
CONF_DIR := conf

GO_FILES := $(shell find . -name "*.go")

HOME_SRC := cmd/home/main.go
HOME_DST := $(BIN_DIR)/home

SITE_SRC := cmd/site/main.go
SITE_DST := $(BIN_DIR)/main

DECRYPT_SRC := cmd/decrypt/main.go
DECRYPT_DST := $(BIN_DIR)/decrypt

LDFLAGS ?= 

@PHONY: run debug build build_template build_save_docker bin_dir clean build_site_alpine_static_for_docker build_home_alpine_static_for_docker setup_tailwind

all: build

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

$(SITE_DST): $(SITE_SRC) $(GO_FILES) | build_template bin_dir
	go build -ldflags "$(LDFLAGS)" -o $(SITE_DST) $(SITE_SRC) 

$(DECRYPT_DST): $(DECRYPT_SRC) $(GO_FILES) | build_template bin_dir
	go build -ldflags "$(LDFLAGS)" -tags=include_private_key -o $(DECRYPT_DST) $(DECRYPT_SRC)

$(HOME_DST): $(HOME_SRC) $(GO_FILES) | build_template bin_dir
	go build -ldflags "$(LDFLAGS)" -tags=include_private_key -o $(HOME_DST) $(HOME_SRC)

build: | $(HOME_DST) $(SITE_DST) $(DECRYPT_DST) build_template bin_dir
	echo "Done"

build_site_alpine_static_for_docker:
	CGO_ENABLED=1 GOOS=linux go build -ldflags "$(LDFLAGS)" -o main -ldflags "-s" $(SITE_SRC)

build_home_alpine_static_for_docker:
	CGO_ENABLED=1 GOOS=linux go build -ldflags "$(LDFLAGS)" -o main -ldflags "-s" $(HOME_SRC)

build_docker: $(SITE_SRC) | build_template bin_dir
	docker build -t "astoryofand:${VERSION}" -f site.Dockerfile .
	docker build -t "astoryofand-home:${VERSION}" -f home.Dockerfile .

build_save_docker: $(SITE_SRC) | build_docker
	docker save -o bin/astoryofand.tar "astoryofand:${VERSION}"

run:  | build_template
	# @DEBUG=false LOG_LEVEL=2 go run $(SITE_SRC)
	@DEBUG=false LOG_LEVEL=2 go run $(HOME_SRC)

debug: | build_template
	gofmt -w -s .
	goimports -w .
	@DEBUG=true go run $(SITE_SRC)

