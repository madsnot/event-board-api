.PHONY: lint
lint:
	golangci-lint run --config ./build/.golangci.yaml ./...

.PHONY: up_local
up_local:
	COMPOSE_PROJECT_NAME=event-board-local docker-compose -f build/docker-compose.yaml up -d --remove-orphans --force-recreate

.PHONY: stop_local
stop_local:
	COMPOSE_PROJECT_NAME=event-board-local docker-compose -f build/docker-compose.yaml down

