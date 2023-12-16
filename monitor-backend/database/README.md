# Migrations

## Tools used

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.14.1
```

## Local postgres

```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=<insert_user> \
  -e POSTGRES_PASSWORD=<insert_password> \
  -e POSTGRES_DB=<insert_db> \
  -p 5432:5432 \
  postgres:15.2-alpine
```
or 

```yaml
version: "3.9"
services:
  testdb:
    container_name: testdb
    image: "postgres:15.2-alpine"
    restart: unless-stopped
    ports:
      - 5432:5432
    environment:
      - POSTGRES_HOST_AUTH_METHOD=trust
      - POSTGRES_USER=<insert_user> 
      - POSTGRES_PASSWORD=<insert_password> 
      - POSTGRES_DB=<insert_db>
    volumes:
      - ./localData:/var/lib/postgresql/data
```

and run 

```
docker compose up -d 
```

## Create and run migrations

run migrations:

```
migrate -path database/migrations/ -database <link to database> up
```

Create migrations:

```
migrate create -ext sql -dir database/migrations/ <migration name>
```
and edit the generated migration files.
