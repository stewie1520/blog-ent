![CI](https://github.com/stewie1520/blog_ent/actions/workflows/go.yml/badge.svg)

## Getting Started 🎃

```sh
docker-compose up -d
cp ./config.example.yaml ./config.yaml
```

Make sure to update config.yaml with corresponding env from `docker-compose.yml`.

## Debug 🐞

```sh
make debug
```

## Watch mode 👀

```sh
make watch
```

## Swagger 🧾
http://localhost:8000/swagger/index.html

## Migration 💿

Please refer to [atlasgo.io](https://atlasgo.io/getting-started/) documentation and [ent migration](https://entgo.io/docs/data-migrations) documentation.
