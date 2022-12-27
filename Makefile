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
	go install github.com/vektra/mockery/v2@latest 1> /dev/null
	go install gotest.tools/gotestsum@latest 1> /dev/null
	go install github.com/boumenot/gocover-cobertura@latest 1> /dev/null
	go install github.com/ggere/gototal-cobertura@latest 1> /dev/null

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

mock:
	rm -rf mocks
	mockery --all --keeptree 
	mockery --all --output mocks/proto --srcpkg github.com/teguh-satriya/privy-go/proto/cakes/v1
	mockery --all --output mocks/package --srcpkg google.golang.org/grpc/grpclog

test: 
	gotestsum --format testname --junitfile junit.xml -- -coverprofile=coverage.lcov.info -covermode count ./...
	gocover-cobertura < coverage.lcov.info > coverage.xml
	gototal-cobertura < coverage.xml