# Go parameters
GOCMD				:= go
PROJECT_PKGS        := $(shell $(GOCMD) list ./... | grep -v '/vendor/')

.PHONY: run test docker-compose-build-api docker-compose-up-api docker-compose-stop-api help all

run: docker-compose-build-api docker-compose-up-api

test:
	@for pkg in $(PROJECT_PKGS); do \
		$(GOCMD) test -v -race -cover $$pkg || exit 1; \
	done

docker-compose-build-api:
	@docker-compose build

docker-compose-up-api:
	@docker-compose up

docker-compose-stop-api:
	@docker-compose stop

all:help

help:
	@echo
	@echo
	@echo
	@echo "Targets are:\n"
	@echo "run"
	@echo " docker build and run"
	@echo
	@echo "test"
	@echo " start unit tests"
	@echo
	@echo
