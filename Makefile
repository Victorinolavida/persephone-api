DB_MIGRATIONS_PATH=resources/migrations
CONFIG_FILE=config.yml

PORT= $(shell yq '.database.port' < $(CONFIG_FILE))
DB_NAME= $(shell yq '.database.dbname' < $(CONFIG_FILE))
DB_USER= $(shell yq '.database.username' < $(CONFIG_FILE))
DB_PASSWORD= $(shell yq '.database.password' < $(CONFIG_FILE))
DB_HOST= $(shell yq '.database.host' < $(CONFIG_FILE))
DSN= postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(PORT)/$(DB_NAME)?sslmode=disable

## help: print this help message
.PHONY: help
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo "Are you sure? [y/N] " && read ans && [ $${ans:-N} = y ]

## run: run the api server
.PHONY: run
run:
	go run ./cmd/main.go
## build: build the api server
.PHONY: build
build:
	go build -o ./bin/main ./cmd/main.go

## test: run tests
.PHONY: test
test:
	go test -v ./...

## tidy: tidy go modules
.PHONY: tidy
tidy:
	go mod tidy

## fmt: format go code
.PHONY: fmt
fmt:
	go fmt ./...

## migration/new name=$1: create a new migration
.PHONY: migration/new
migration/new:
	migrate create -seq -ext=.sql -dir=./$(DB_MIGRATIONS_PATH) ${name}

## migration/up: apply all migrations
.PHONY: migration/up
migration/up: confirm
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DSN) up


## migration/down: rollback all migrations
.PHONY: migration/down
migration/down:
	migrate -path $(DB_MIGRATIONS_PATH) -database $(DSN) down