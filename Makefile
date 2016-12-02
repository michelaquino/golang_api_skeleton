# Go parameters
GOCMD				:= go
GOPATH 				:= $(shell go env GOPATH)
PROJECT_PKGS        := $(shell $(GOCMD) list ./... | grep -v '/vendor/')

.PHONY: run
run: 
	$(GOCMD) run -race main.go

.PHONY: test
test: setup
	@for pkg in $(PROJECT_PKGS); do \
		$(GOCMD) test -v -race $$pkg || exit 1; \
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