
# User Service

## Requirement
- Golang 1.23
- Makefile
- PostgreSQL 

## Quick start

```sh
# Install required tools  
$ make setup

# Generate API documentation
$ make doc

# Create configuration files and update the configs
$ cp configs/config.yml.example config.yml

# Run app from source
$ make run

# Generate binary file
$ make build

# Run app auto rebuild
$ make watch

# Run app in docker
$ make docker-run

# Run unit test
$ make test

# Create new migration file
$ make db-migration name=create_table

# Update database 
$ make db-migrate

# Rollback database 
$ make db-rollback

# Run app binary 
$ ./bin/runner serve

```

## Features
- [x] API Documentation
- [x] Unit Testing
- [x] Dockerized
- [x] DB Migration
- [x] Service Metric
- [ ] Log Wrapper
- [ ] GRPC / HTTP communication
- [ ] Kafka integration
- [ ] Redis integration
