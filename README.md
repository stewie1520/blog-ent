![CI](https://github.com/stewie1520/blog/actions/workflows/go.yml/badge.svg)

## Getting Started 🎃

```sh
docker-compose up -d
cp ./config.example.yaml ./config.yaml
```

Make sure to update config.yaml with corresponding env from `docker-compose.yml`.

## Debug

```sh
make debug
```

## Watch mode

```sh
make watch
```

## Migration

To create a migration file, run
```sh
make migrate-gen name=<migration_file_name>
```

To run migration
```sh
make migrate-up database="<database_url>"
```

To undo a migration
```sh
make migrate-down database="<database_url>" step=1
```

We're using [golang-migrate](github.com/golang-migrate/migrate) underlying, please go check their documentation.
