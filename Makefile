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

start-dev-env: ## start dev env
	@docker-compose -f .docker/dev/docker-compose.yml up -d

stop-dev-env: ## stop dev env
	@docker-compose -f .docker/dev/docker-compose.yml down

post: ## post [DATA].json file
	@curl -X POST -H "Content-Type: application/json" -d @./tmp/post/${DATA}.json http://127.0.0.1:8080/${DATA}

patch: ## patch [ID] with [DATA].json file
	@curl -X PATCH -H "Content-Type: application/json" -d @./tmp/patch/${DATA}.json http://127.0.0.1:8080/${DATA}/${ID}

get: ## get [DATA]
	@curl -X GET http://127.0.0.1:8080/${DATA}

