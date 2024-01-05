run:
	@templ generate
	@go run cmd/itcontainers/main.go

migrate:
	@go run cmd/migrations/migrate.go

build:
	@templ generate
	@go build cmd/itcontainers/main.go

watch:
	@CompileDaemon -command="./main" -exclude-dir=.git -build="make build" -exclude="*_templ.go" -include="*.templ"
