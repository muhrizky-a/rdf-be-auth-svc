# This Project contain this package

## Golang

go version -> go1.24.1

### Testify

go get github.com/stretchr/testify

### Assert

go get github.com/stretchr/testify/assert

### Mock

go get github.com/stretchr/testify/mock

### GRPC

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

export PATH="$PATH:$(go env GOPATH)/bin"

## Gorm

go get gorm.io/gorm
go get gorm.io/driver/postgres

## PostgreSQL

docker image -> postgres:17-alpine

## Go-migrate

curl https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.windows-amd64.zip -o migrate.windows-amd64.zip

or

wget https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.windows-amd64.zip

# Tutorials

## Load env

## Load env in PowerShell

$env:DB_HOST="localhost";
$env:DB_PORT="5432";
$env:DB_USER="developer";
$env:DB_PASS="supersecretpassword";
$env:DB_NAME="rdf_auth_db_test";
$env:DB_ENABLE_SSL="disable";

## Database migration

### Create migration

<!-- migrate create -ext sql -dir db/migrations/ {{.name}} -tz UTC -->

migrate create -ext sql -dir db/migrations/ create-table-scopes -tz UTC

### Migrate up

<!-- migrate -path db/migrations -database"postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" up -->

migrate -path db/migrations -database "postgresql://developer:supersecretpassword@localhost:5432/rdf_auth_db_test?sslmode=disable" up

### Migrate down

<!-- migrate -path db/migrations -database"postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" down -->

migrate -path db/migrations -database "postgresql://developer:supersecretpassword@localhost:5432/rdf_auth_db_test?sslmode=disable" down

### Migrate down all

<!-- migrate -path db/migrations -database"postgresql://username:secretkey@localhost:5432/database_name?sslmode=disable" down  -all -->

migrate -path db/migrations -database "postgresql://developer:supersecretpassword@localhost:5432/rdf_auth_db_test?sslmode=disable" down -all

### Seeds DB

go run db/seeds/main.go --table=scopes

## Testing

### Execute testing

<!-- go test -count=1 -timeout 30m -p 1  <module/package> -->

go test -count=1 -timeout 30m -p 1 github.com/ryakadev/rdf-be-auth-svc/repository
go test -count=1 -timeout 30m -p 1 github.com/ryakadev/rdf-be-auth-svc/usecase
go test -count=1 -timeout 30m -p 1 github.com/ryakadev/rdf-be-auth-svc/delivery/grpc
