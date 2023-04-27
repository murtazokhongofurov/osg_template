CURRENT_DIR=$(shell pwd)

-include .env

DB_URL="postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable"

migrate_file:
	migrate create -ext sql -dir migrations -seq create_commit_table


migrate_up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate_down:
	migrate -path migrations -database "$(DB_URL)" -verbose down 

migrate_forse:
	migrate -path migrations -database "$(DB_URL)" -verbose forse 

run:
	go run cmd/main.go