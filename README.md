
# ***Hello Dear Friend, its HaveIdea!***

## run
### `go run cmd/main.go`

## migration create
### `migrate create -ext sql -dir db/migrations -seq ${<migration_name>}`

## migration up
### `migrate -database "postgres://postgres:<pg_password>@localhost:5432/invest_db?sslmode=disable" -path db/migrations up`

## migration down
### `migrate -database "postgres://postgres:<pg_password>@localhost:5432/invest_db?sslmode=disable" -path db/migrations down`

## see jwt secret
### `printf "\nJWT_SECRET=\"$(shell openssl rand -base64 32)\"\n" >> .env`