run:
	@go run cmd/main.go

build:
	@go build cmd/main.go

watch:
	@CompileDaemon -command="./main" -exclude-dir=.git -build="make build"

migrate:
	@goose -dir db/migrations sqlite3 ./data/db.sqlite3 up

migrate-down:
	@goose -dir db/migrations sqlite3 ./data/db.sqlite3 down

migrate-reset:
	@goose -dir db/migrations sqlite3 ./data/db.sqlite3 reset
