•⁠  ⁠Postgress connection - OK
•⁠  ⁠ENV variables - OK
•⁠  ⁠CRUD Albums - OK
•⁠  ⁠Migrations - OK
•⁠  ⁠Dockerize - OK
•⁠  ⁠Unit tests - Maso
•⁠  ⁠Swagger -- Mañana
create example .env

Nice to Have
•⁠  ⁠Middleare example - OK
•⁠  ⁠Case insensitive
•⁠  ⁠Configure hot reload
•⁠  ⁠Call an external service
•⁠  ⁠Rate limit for that external service
•⁠  ⁠Linter
•⁠  ⁠Put an image into S3 or something like that


Others
Update my resumee in .pdf
Create GitHub page
Create CLI


THIS WORKS

up

docker run --rm -v $(pwd)/db/migrations:/migrations --network host migrate/migrate -database "postgres://postgres:p0stgr3s.4dm1n.p0wer@0.0.0.0:5432/dev?sslmode=disable" -path /migrations up

DOWN

docker run --rm -v $(pwd)/db/migrations:/migrations --network host migrate/migrate -database "postgres://postgres:p0stgr3s.4dm1n.p0wer@0.0.0.0:5432/dev?sslmode=disable" -path /migrations down -all