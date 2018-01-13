# Go parameters
GOCMD				:= go
PROJECT_PKGS        := $(shell $(GOCMD) list ./... | grep -v '/vendor/')

.PHONY: all help run test docker-compose-build-api docker-compose-up-api docker-compose-stop-api

all: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test:	## start unit tests
	@for pkg in $(PROJECT_PKGS); do \
		$(GOCMD) test -v -race -cover $$pkg || exit 1; \
	done

docker-compose-build-api: ## Build the application image
	@docker-compose build

docker-compose-up-api: ## Up the application image
	@docker-compose up

docker-compose-stop-api: ## Stop the application containers
	@docker-compose stop

run: docker-compose-build-api docker-compose-up-api ## Run the application loccaly