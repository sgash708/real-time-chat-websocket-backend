.PHONY: up
up:
	@make down
	@docker-compose up

.PHONY: upd
upd:
	@docker-compose up -d

.PHONY: down
down:
	@docker-compose down

.PHONY: exec
exec:
	@docker-compose exec rt-chat-service sh

.PHONY: build
build:
	@docker-compose build --no-cache

.PHONY: build-up
build-up:
	@docker-compose up -d --remove-orphans

.PHONY: img-prune
img-prune:
	@docker volume ls | xargs docker volume rm | 2>/dev/null
	@docker image prune -a
	@docker builder prune
