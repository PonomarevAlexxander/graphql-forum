.PHONY: build
build: ## Build app
	go build -o ./bin/forum ./cmd/app/main.go

.PHONY: unit-test
unit-test: ## Run unit tests
	go test ./...

.PHONY: install-tools
install-tools: ## Add some dependencies for generation and migrations
	go get -u github.com/99designs/gqlgen@latest
	go get -u github.com/vektah/dataloaden@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: migration
migration: ## Run migrations
	goose -dir ./migrations postgres "host=localhost user=pguser password=pgpwd dbname=forum sslmode=disable" up

.PHONY: gql-gen
gql-gen: ## Generate GQL models and resolvers
	go run github.com/99designs/gqlgen generate

.PHONY: gogen
gogen: ## Generate all go generate fragments
	go generate ./...

version:=
.PHONY: docker-build
docker-build: ## Build docker image
	docker build -t gql-forum .
ifdef version
	docker image tag gql-forum:latest gql-forum:$(version)
endif

.PHONY: start-dev
start-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		-f deployments/docker/app.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		up -d

args=
.PHONY: stop-dev
stop-dev:
	docker-compose \
		-f deployments/docker/postgres.docker-compose.yaml \
		-f deployments/docker/app.docker-compose.yaml \
		--env-file=deployments/docker/.env \
		down $(args)

file:=
CONFIG_DIR:=config/app/config.yaml
.PHONY: docker-run
docker-run: ## Run docker image
	docker run -p 8080:8080 -v "$(dir $(realpath $(lastword $(MAKEFILE_LIST))))${CONFIG_DIR}:/${CONFIG_DIR}" gql-forum
