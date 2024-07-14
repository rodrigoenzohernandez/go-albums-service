# Makefile

run-local:
	docker-compose up

build:
	docker-compose build

migrate-up:
	docker-compose run --rm migrate-up

migrate-down:
	docker-compose run --rm migrate-down
