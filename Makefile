build:
	@echo "Building..."
	@go build -o bin/main main.go

run:
	@go run cmd/console/main.go

migration:
	@go run cmd/migration/main.go

db-up:
	@docker compose up -d

db-down:
	@docker compose down

db-start:
	@docker compose start

db-stop:
	@docker compose stop
