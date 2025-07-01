# Interview Test Ava Airline

## CGO!
May you get error cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in $PATH while running the application, please install gcc.

## Dependencies
you need to install the following dependencies:
- gcc
- go v1.24+
- sqlc `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

```
you don't need any database servers to run the application. it's configured to work with sqlite which is file db working properly in local.
migrations will automatically applies the changes to the database whenever you run the application.
```
```
test cases will start webserver and database in memory, so if you have run the application before running the tests, please stop it first.
```

## Running tests
```sh
go test -v ./e2e
```

## Running application
```sh
cd ./app
go run ../
```

## Running dockerized application
```sh
docker build -t interview .
docker run -p 8080:8080 -d interview:latest
```

## Making some changes
- apis: define them in the `api` directory.
- models: define the migrations, query and schema in the `database` directory then generate the code using `sqlc generate` command.
- end-to-end tests: define them in the `e2e` directory.
- unit tests: define them in the `unit's` directory.
