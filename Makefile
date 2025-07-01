VERSION:=$(shell grep 'VERSION' internal/constants/version.go | awk '{ print $$4 }' | tr -d '"')

.DEFAULT_GOAL := lint

# build
build:
	@go build -o ./bin/runner ./cmd/server.go

clean:
	@echo "Cleaning..."
	@go mod tidy
	@rm ./bin/runner

setup: setup-migration setup-doc setup-test setup-air

# Run the application
run:
	@go run cmd/server.go serve

# Create DB container
docker-run:
	@docker compose up --build

# Shutdown DB container
docker-down:
	@docker compose down
	
# # Start mysql container
# mysql_install:
# 	docker run -itd --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=$(DBSECRET) mysql

# # Create database
# db_create:
# 	docker exec -it mysql-test mysql -p$(DBSECRET) -e "create database $(DBNAME)"

# Generate new migration
setup-migration:
	@go install -tags='no_mysql no_sqlite3 no_ydb' github.com/pressly/goose/v3/cmd/goose@latest
	
# make migration name=update_table_user
db-migration:
	@goose create $(name) sql

db-status:
	@goose status

db-migrate:
	@echo "Migrating..."
	@goose up

db-rollback:
	@echo "Rollback..."
	@goose down

db-seed:
	@go run cmd/main.go seed --count $(count) --name $(name)

# lint source file
lint:
	@golangci-lint run

# unit test coverage
setup-test:
	@go install github.com/vektra/mockery/v3@v3.2.5

test:
	@echo "Running test, please wait ..."
	@go test ./internal/... --covermode=count '-gcflags=all=-N -l' -v

# documentation
setup-doc:
	@go install github.com/swaggo/swag/cmd/swag@latest
	
doc-fmt:
	@swag fmt
	
doc:
	@swag init -g ./internal/server/server.go 

setup-air:
	@go install github.com/cosmtrek/air@latest

watch:
	@air -c .air.toml

.PHONY: mysql_install db_create lint
