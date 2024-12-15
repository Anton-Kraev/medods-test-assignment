.PHONY:
.SILENT:

all: postgres migrate app

postgres:
	docker-compose up --build -d postgres

migrate:
	./scripts/migrate.sh

app:
	docker-compose up --build -d app
