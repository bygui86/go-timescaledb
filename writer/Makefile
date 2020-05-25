
# VARIABLES
NAME = timescaledb-writer
CONTAINER_NAME = samples/$(NAME)


# ENVIRONMENT VARIABLES
export GO111MODULE = on


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

start-timescaledb :		## Run TimescaleDB in a container
	docker run -d --rm --name timescaledb \
		-e POSTGRES_PASSWORD=supersecret \
		-p 5432:5432 \
		timescale/timescaledb:latest-pg12

stop-timescaledb :		## Stop TimescaleDB container
	docker stop timescaledb

connect-timescaledb :		## Connect to TimescaleDB container
	docker exec -it timescaledb bash

## application

build :		## Build application
	go build

start :		## Run application from source code
	godotenv -f local.env go run main.go

start-binary : build		## Run application from binary
	source local-binary.env && ./writer

## cointaier

container-build :		## Build container image
	docker build . -t $(CONTAINER_NAME)

## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
