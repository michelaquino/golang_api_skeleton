include Makefile.vars

.PHONY: run
run: docker-compose-build-api docker-compose-up-api

.PHONY: test
test:
	@for pkg in $(PROJECT_PKGS); do \
		$(GOCMD) test -v -race -cover $$pkg || exit 1; \
	done

.PHONY: docker-compose-build-api
docker-compose-build-api:
	@docker-compose build

.PHONY: docker-compose-up-api
docker-compose-up-api:
	@docker-compose up

.PHONY: docker-compose-stop-api
docker-compose-stop-api:
	@docker-compose stop

.PHONY: all
all:help

.PHONY: help
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
