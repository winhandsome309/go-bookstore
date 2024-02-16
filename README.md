# Go Bookstore

A bookstore e-commerce websites

## Tech stack

- gin-gonic
- swagger
- jwt-go
- logrus (logging)
- crypto/bcrypt (encrypt)
- air (live reload)
- gorm
- postgreSQL (database)

## Config

- Path: `pkg/config/config.yaml`

```
http_port: 8080
auth_secret: 000000
database_uri: host=localhost user=postgres password=30092002 dbname=bookstore port=5432 sslmode=disable=
```

## Run

- 2 options

```shell script
$ go run cmd/api/main.go
```

```shell script
$ air
```

## Document

- API document at: `http://localhost:8888/swagger/index.html`

## What's next

- Using gRPC, message queue...
