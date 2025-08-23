include .env
export $(shell sed 's/=.*//' .env)

MIGRATE_DIR=database/migrations

DB_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable


create_postgres_db:
	docker run --name $(POSTGRES_CONTAINER_NAME) -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -p $(POSTGRES_PORT):5432 -d postgres:14

create_migration:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(NAME)

migrate_up:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) up

migrate_down:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) down

migrate_clean:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) force 1

migrate_status:
	migrate -database "$(DB_URL)" -path $(MIGRATE_DIR) version