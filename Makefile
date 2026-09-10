include .env
export

export PROJECT_ROOT=$(shell pwd)

pg-up:
	@docker compose up -d todoapp-postgres

pg-down:
	@docker compose down todoapp-postgres

volume-clean:
	@read -p "Очистить все volume файлы pgdata? [y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down todoapp-postgres && \
		docker run --rm -v "$(PROJECT_ROOT)/out:/out" alpine sh -c "rm -rf /out/pgdata" && \
		echo "Volume файлы очищены"; \
	else \
		echo "Очистка volume файлов отменена"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсуствует нужный параметр seq. Пример вызова Makefile: make migrate-create seq=init"; \
		exit 1; \
	fi;

	docker compose run --rm todoapp-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Отсуствует нужный параметр action. Пример вызова Makefile: make migrate-action action=up"; \
		exit 1; \
	fi;
	docker compose run --rm todoapp-postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"