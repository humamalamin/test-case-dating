# SIMPLE API DATING

## Description

This project was created for technical test. I'm try to implement clean code my version and unit testing

## Tech Stack
- Golang : https://github.com/golang/go
- Mysql (Database) : https://github.com/mysql/mysql-server
- Docker (Container) : https://www.docker.com/get-started/
- Docker Compose : https://docs.docker.com/compose/

## Framework & Library
- Gorilla Mux (HTTP Framework) : https://github.com/gorilla/mux
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- JWT (Token Authentication) : https://github.com/golang-jwt/jwt

## Configuration

All configuration is in `.env` file.
If use docker-compose, please see config DB inside file docker-compose.

## Database Migration

All database migration is in `db/` folder.

### Initial Project

```shell
make init
```

### Create Migration

```shell
migrate create -ext sql db/ create_table_xxx
```

### Run Migration

```shell
migrate -database "mysql://<user>:<password>@tcp(<host>:<port>)/<database>" -path /db up
```

### Run Coverage Test

```shell
make test-coverage:
```

### Run Generate Mock File

```shell
make generate-mock
```

## Run Application

### Run With Docker Compose

```bash
make build-docker
```

### Run web server

```bash
make run

or

docker-compose up --build -d
```

## Deployment Manual

### Spesific minimum server
- RAM up 2GB
- VPS
- Storage Up 20GB
- CPU 2 Core
- OS: Ubuntu

### Requirement Server
- Golang : https://github.com/golang/go
- Mysql (Database) : https://github.com/mysql/mysql-server
- Docker (Container) : https://www.docker.com/get-started/
- Docker Compose : https://docs.docker.com/compose/

### Run Application in Server

- All requirement server installed.
- Create new database Mysql
- following step by step in below:

```bash
git clone git@github.com:humamalamin/test-case-dating.git

cd test-case-dating

go mod download

cp .env.example to .env

config file .env

docker-compose up --build -d
```
