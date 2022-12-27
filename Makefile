MAKEFLAGS += --silent

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

DATABASE_URL="mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}?parseTime=true"

prepare-migration:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz
	mv migrate /usr/bin &>/dev/null

create-migration:
	$(eval timestamp := $(shell date +%s))
	touch db/migrations/$(timestamp)_${name}.up.sql
	touch db/migrations/$(timestamp)_${name}.down.sql

rollback-migration:
	migrate --path=db/migrations/ \
			--database ${DATABASE_URL} down

run-migration:
	migrate --path=db/migrations/ \
			--database ${DATABASE_URL} up

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