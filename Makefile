.PHONY: postgres create-migrate migrate

postgres:
	docker-compose up -d postgres

create-migrate:
	migrate create -ext postgresql -dir ./postgres/migrations -seq <name>

migrate:
	docker-compose run migrate