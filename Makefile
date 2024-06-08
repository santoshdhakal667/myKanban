build:
	@go build -o ./bin/backend .

compose:
	@docker compose up -d

dev:	compose
	@air 

migration_up: compose 
	migrate -path database/migration/ -database "mysql://kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true" -verbose up

migration_down: compose
	migrate -path database/migration/ -database "mysql://kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true" -verbose down

migration_fix: compose
	migrate -path database/migration/ -database "mysql://kanban:bw1qJGj@tcp(127.0.0.1:6000)/kanban?multiStatements=true" force VERSION

backend: build
	@./bin/backend

generate-bindata:
	@go-bindata -o ./database/db.go -prefix "./database" -pkg init.sql ./database
	
