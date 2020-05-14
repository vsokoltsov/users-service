include .env

.PHONY: up
up: docker-compose.yaml
	@echo "$@"
	docker-compose up

.PHONY: down
down: docker-compose.yaml
	@echo "$@"
	docker-compose stop

.PHONY: clean
clean:
	@echo "$@"
	docker kill $(shell docker ps -a -q) || true
	docker rm $(shell docker ps -a -q) || true

.PHONY: shell
shell:
	@echo "$@"
	docker exec -it users_service \
		/bin/sh

.PHONY: test
test:
ifndef $(ARGS)
	@echo 'no ARGS around'
	$(eval ARGS := "./...")
endif
	docker exec -it users_db /bin/sh -c \
		"dropdb --if-exists -U "${POSTGRES_USER}" "${POSTGRES_DB_TEST}" && createdb -U "${POSTGRES_USER}" "${POSTGRES_DB_TEST}""
	docker exec -it -e APP_ENV=test -e POSTGRES_DB="${POSTGRES_DB_TEST}" users_service \
			go test -v "${ARGS}"


.PHONY: add_migration
add_migration:
	@echo "$@"
	docker exec -it users_service /bin/sh -c \
		"goose -dir /app/app/migrations/ postgres \"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}\" create ${ARGS} sql"

.PHONY: migrate_up
migrate_up:
	@echo "$@"
	docker exec -it users_service /bin/sh -c \
		"goose -dir /app/app/migrations/ postgres \"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}\" up"

.PHONY: migrate_down
migrate_down:
	@echo "$@"
	docker exec -it users_service /bin/sh -c \
		"goose -dir /app/app/migrations/ postgres \"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:5432/${POSTGRES_DB}?sslmode=${POSTGRES_SSL}\" down"

.PHONY: generate_proto
generate_proto:
	@echo "$@"
	docker exec -it users_service /bin/sh -c \
		"protoc --go_out=plugins=grpc:/app/app/grpc/ $(ARGS)"