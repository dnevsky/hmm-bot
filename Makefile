.PHONY: postgres create-migrate migrate build run-bot run shutdown

postgres:
	docker-compose up -d postgres

mysql:
	docker-compose up -d mysql

create-migrate:
	migrate create -ext sql -dir ./postgres/migrations -seq <name>

migrate:
	docker-compose run migrate

build:
	docker build --tag dnevsky/hmm-bot .

run:
	docker-compose up -d hmm-bot

shutdown:
	docker-compose down