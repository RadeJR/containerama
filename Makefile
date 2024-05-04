run:
	@templ generate
	@go run cmd/containerama/main.go

build:
	@templ generate
	@go build cmd/containerama/main.go

watch:
	@CompileDaemon -command="./main" -exclude-dir=.git -build="make build" -exclude="*_templ.go" -include="*.templ"

migrate:
	@goose -dir db/migrations sqlite3 ./db.sqlite3 up

migrate-down:
	@goose -dir db/migrations sqlite3 ./db.sqlite3 down
