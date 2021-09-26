# first_go_app

## Dependencies
 - go
 - a mysql/mariadb database

```bash
go mod init
```

## Edit config.json
```json
{
  "server": {
    "host": "0.0.0.0",
    "port": "8080"
  },
  "database": {
    "driver": "mysql",
    "username": "dbuser",
    "password": "password",
    "host": "127.0.0.1",
    "port": "3306",
    "name": "dbname"
  },
  "logger": {
    "debug": true,
    "logfile_path": "./log/",
    "logfile_name": "first_go_app.log"
  }
}
```

## Setup Database

```bash
go run main.go migrate
```

## Run it

```bash
go run main.go serve
```

## Create test items

```bash
curl -X POST localhost:8080/tests -H 'Content-Type: application/json' -d '{"Name":"testname","Description":"testdescription"}'
```

## Help

```bash
go run main.go -h
```