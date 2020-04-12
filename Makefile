include .env

.PHONY: up
up: docker-compose.yaml
	@echo "$@"
	docker-compose up

.PHONY: down
down: docker-compose.yaml
	@echo "$@"
	docker-compose down

.PHONY: clean
clean:
	@echo "$@"
	docker kill $(shell docker ps -a -q) || true
	docker rm $(shell docker ps -a -q) || true

.PHONY: shell
shell:
	@echo "$@"
	docker exec -it users_service \
		/bin/bash

.PHONY: test
test:
ifndef $(ARGS)
	@echo 'no ARGS around'
	$(eval ARGS := "./...")
endif
	docker exec -it users_db /bin/bash -c \
		"dropdb --if-exists -U "${POSTGRES_USER}" "${POSTGRES_DB_TEST}" && createdb -U "${POSTGRES_USER}" "${POSTGRES_DB_TEST}""
	docker exec -it -e APP_ENV=test -e POSTGRES_DB="${POSTGRES_DB_TEST}" users_service \
			go test -v "${ARGS}"