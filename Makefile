# Go parameters
GOCMD				:= go
PROJECT_PKGS        := $(shell $(GOCMD) list ./... | grep -v '/vendor/')

.PHONY: run
run: docker-compose-build-api docker-compose-up-api

.PHONY: test
test:
	@for pkg in $(PROJECT_PKGS); do \
		$(GOCMD) test -v -race -cover $$pkg || exit 1; \
	done

######## Docker compose ########
.PHONY:	docker-compose-build-api
docker-compose-build-api:
	docker-compose build

.PHONY:	docker-compose-up-api
docker-compose-up-api:
	docker-compose up

.PHONY:	docker-compose-stop-api
docker-compose-stop-api:
	docker-compose stop