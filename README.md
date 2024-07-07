
web-service-gin

API that provides access to albums. All data is stored in memory.

# Endpoints

## /albums

GET – Get a list of all albums, returned as JSON.

```bash
curl http://localhost:8080/albums
```

POST – Add a new album from request data sent as JSON.

```bash
curl http://localhost:8080/albums \
    --include --header \
    "Content-Type: application/json" \
    --request "POST" --data \
    '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

## /albums/:id

GET – Get an album by its ID, returning the album data as JSON.

```bash
curl http://localhost:8080/albums/2
```

# How to start the project?

Install the dependencies
```shell
go get .
```

Run
```shell
go run .
```

