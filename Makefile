# Docker image version
DK_REGISTRY ?= jbustin1/

DK_VERSION ?= 0.0.1
DK_NAME ?= go-mservice-bench
DK_IMAGE ?= $(DK_REGISTRY)$(DK_NAME):$(DK_VERSION)
DK_LATEST ?= $(DK_REGISTRY)$(DK_NAME):latest

.DEFAULT_GOAL := help

help: ## Display this help
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 }'

start-redis: ## start dev env
	@docker-compose -f .docker/docker-compose.yml up -d

stop-redis: ## stop dev env
	@docker-compose -f .docker/docker-compose.yml down

post: ## post [DATA].json file
	@curl -X POST -H "Content-Type: application/json" -d @./samples/post/${DATA}.json http://127.0.0.1:8080/${DATA}

patch: ## patch [ID] with [DATA].json file
	@curl -X PATCH -H "Content-Type: application/json" -d @./samples/patch/${DATA}.json http://127.0.0.1:8080/${DATA}/${ID}

get-all: ## get [DATA]
	@curl -X GET http://127.0.0.1:8080/${DATA}

get: ## get [DATA] [ID]
	@curl -X GET http://127.0.0.1:8080/${DATA}/${ID}

delete: ## delete [ID]
	@curl -X DELETE http://127.0.0.1:8080/${DATA}/${ID}

bench: ## bench [PACKAGE] / [TEST]
	go test -benchmem -run=^$$ -bench ^${TEST}$$  github.com/go-mservice-bench/lib/${PACKAGE}

attack: ## start [COUNT] * [CONCURRENT] transactions
	@for run in {1..${COUNT}}; do \
		for run in {1..${CONCURRENT}}; do \
			DATA=transaction make post & \
		done \
	done

