# Makefile

# Run local
local:
	go run cmd/app/main.go

install:
	go mod tidy

build:
	go build -o bin/app cmd/app/main.go

migrate-up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations -verbose up

migrate-down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations -verbose up