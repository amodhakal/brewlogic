# Use the .env file for the data
ifneq (,$(wildcard .env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

PSQL_STRING := postgres://$(PSQL_USER):$(PSQL_PASSWORD)@$(PSQL_HOST):$(PSQL_PORT)/$(PSQL_DB)

run:
	go run ./src
migrate:
	psql-17 ${PSQL_STRING} -f ./sql/schema.sql
