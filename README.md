# HOTEL-PARSER

This is Golang-program that parses JSON, CSV and XML files containing hotel data and stores this data in the database.
WARNING: checking for the presence of similar data in the database is not performed, the data is simply added so far.

## Firts steps

1. Create and edit `.env` file by `.env.example`
2. Edit `docker-compose.yaml` file:
    - set volume binding to the dir with raw JSON, CSV and XML files
    - set volume binding to the `.env` file
    - set necessary ports binding
3. Run `docker-compose up -d`

## Migrate DB

See <https://github.com/golang-migrate/migrate/>

1. Install:

    ```bash
    go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
    ```

2. Create migration files (up and down):

    ```bash
    migrate create -ext sql -dir db/migrations/ -seq MIGRATION_NAME
    ```

3. Write migration code to VERSION.MIGRATION_NAME.up.sql and VERSION.MIGRATION_NAME.down.sql files.

4. Run migration:

    4.1 directly:

    ```bash
    migrate -database postgres://user:password@host:port/name?sslmode=disable -path db/migrations up
    ```

    or

    ```bash
    migrate -database postgres://user:password@host:port/name?sslmode=disable -path db/migrations up 1
    ```

    or

    ```bash
    migrate -database postgres://user:password@host:port/name?sslmode=disable -path db/migrations down 1
    ```

    4.2 run `docker-compose up` to up all migrations

5. If an error occurs during migration, to force vesion run:

    ```bash
    migrate -database postgres://user:password@host:port/name?sslmode=disable -path db/migrations force LAST_SUCCESSFULL_VERSION
    ```
