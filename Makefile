MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

setup: 
	go mod tidy
	go mod vendor 

develop: stop
	docker-compose up -d 1> /dev/null
	docker-compose logs -f privy-cake-httpd

stop:
	docker-compose stop privy-cake-httpd &> /dev/null

build:
	set -e mkdir target/bin
	go build -o target/bin/privy-cake main.go

install: build
	cp target/bin/* /usr/local/bin/