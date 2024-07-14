# go-albums-service

RestAPI that provides a CRUD of Albums.

# 1. Built With

- [Go](https://golang.org/) - The Go Programming Language
- [Docker](https://www.docker.com/) - Containerization platform
- [Gin](https://github.com/gin-gonic/gin) - A HTTP web framework written in Go
- [PostgreSQL](https://www.postgresql.org/) - Open source object-relational database system
  http://localhost:8081/

# 2. How to start the project?

## Start the containers

```shell
make run-local
```

## Run the migrations

```shell
make migrate-up
```

# 3. Documentation

## Swagger

### Run the command

```shell
make swagger
```

### Open swagger at your web browser

[Swagger docs](http://localhost:8081/)

# 4. Test the endpoints

## GET /albums

Get a list of all albums, returned as JSON.

```bash
curl http://localhost:8080/albums
```

## POST /albums

Adds a new album from request data sent as JSON.

```bash
curl -X 'POST' \
  'http://localhost:8080/albums' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "title": "Blue Train",
  "artist": "John Coltrane",
  "price": 56.99
}'
```

## GET /albums/:id

Gets an album by its ID, returning the album data as JSON.

```bash
curl -X 'GET' http://localhost:8080/albums/0d0d9682-a4c4-47ae-854c-49aaa6f65528
```

## PUT /albums/:id

Updates an album.

```bash
curl -X 'PUT' \
  'http://localhost:8080/albums/0d0d9682-a4c4-47ae-854c-49aaa6f65528' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "title": "8 Mile",
  "artist": "Eminem",
  "price": 56.99
}'
```

## DELETE /albums/:id

Deletes an album.

```bash
curl -X 'DELETE' \
  'http://localhost:8080/albums/0d0d9682-a4c4-47ae-854c-49aaa6f65528' \
  -H 'accept: application/json'
```

# 5. Other commands

## Run the tests

```bash
make test
```

## Delete the resources created on the migration

```bash
make migrate-down
```

## Build docker without executing

```bash
make build
```
