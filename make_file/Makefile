include .env
export $(shell sed 's/=.*//' .env)
run:
	go run main.go

create_db:
	docker run --name tutorial-go -e POSTGRES_USER=$(POSTGRES_USER) -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_DB=$(POSTGRES_DB) -p $(POSTGRES_PORT):5432 -d postgres:14