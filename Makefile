
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## infra

### jaeger

start-jaeger :		## Run Jaeger (all-in-one) in a container
	docker run -d --rm --name jaeger \
		-p 5775:5775/udp \
		-p 5778:5578 \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 9411:9411 \
		-p 14268:14268 \
		-p 14269:14269 \
		-p 16686:16686 \
		jaegertracing/all-in-one

stop-jaeger :		## Stop Jaeger (all-in-one) in container
	docker stop jaeger

open-jaeger-ui :		## Open Jaeger UI in browser
	open http://localhost:16686

### zipkin

start-zipkin :		## Run Zipkin in a container
	docker run -d --rm --name zipkin \
		-p 9411:9411 \
		openzipkin/zipkin

stop-zipkin :		## Stop Zipkin in container
	docker stop zipkin

open-zipkin-ui :		## Open Zipkin UI in browser
	open http://localhost:9411

### postgres

start-timescaledb :		## Run TimescaleDB in a container
	cd writer/ && make start-timescaledb

stop-timescaledb :		## Stop TimescaleDB in container
	cd writer/ && make stop-timescaledb

## applications

start-writer :		## Run writer
	cd writer/ && make start

start-reader :		## Run reader
	cd reader/ && make start

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
