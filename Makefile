run:
	@templ generate
	@go run cmd/containerama/main.go

migrate:
	@go run cmd/migrations/migrate.go

build:
	@templ generate
	@go build cmd/containerama/main.go

watch:
	@CompileDaemon -command="./main" -exclude-dir=.git -build="make build" -exclude="*_templ.go" -include="*.templ"

database:
	@docker compose up -d
