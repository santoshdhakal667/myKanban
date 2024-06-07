build:
	@go build -o ./bin/backend .

dev:
	@air

migrate:
	@

backend:
	@./bin/backend

generate-bindata:
	@go-bindata -o ./database/db.go -prefix "./database" -pkg init.sql ./database
	
