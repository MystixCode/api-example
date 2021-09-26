# first_go_app

## Get dependencies
```bash
go mod init
```

## Run it

```bash
go run main.go serve
```

## Setup Database

```bash
go run main.go migrate
```

## Create test items

```bash
curl -X POST localhost:8080/tests -H 'Content-Type: application/json' -d '{"Name":"testname","Description":"testdescription"}'
```

