build:
	@go build -o bin/invest_app cmd/app/main.go

run:build
	@./bin/invest_app

new_migrate:
	@migrate create -ext sql -dir db/migrations -seq ${name}

migrate-up:
	@migrate -database "postgres://postgres:40029051@localhost:5432/invest_db?sslmode=disable" -path db/migrations up

migrate-down:
	@migrate -database "postgres://postgres:40029051@localhost:5432/invest_db?sslmode=disable" -path db/migrations down 1

jwt_secret:
	@printf "\nJWT_SECRET=\"$(shell openssl rand -base64 32)\"\n" >> .env

