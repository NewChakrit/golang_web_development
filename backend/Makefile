include .env

MIGRATION_PATH = ./db/migrations

.PHONY: migrate-create
migreate-create:
	@migrate create -seq -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_PATH) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATION_PATH) -database=$(DB_PATH) down $(filter-out $@,$(MAKECMDGOALS))