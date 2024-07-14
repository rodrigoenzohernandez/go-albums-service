# Makefile

run-local:
	docker-compose up

test:
	go test ./...

build:
	docker-compose build

migrate-up:
	docker-compose run --rm migrate-up

migrate-down:
	docker-compose run --rm migrate-down

swagger:
	docker run -it -p 8081:8080 -e SWAGGER_JSON=/swagger.json -v ./swagger.json:/swagger.json swaggerapi/swagger-ui
