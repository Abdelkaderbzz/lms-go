.PHONY: help build run test clean docker-up docker-down deps fmt lint seed test-coverage

help:
	@echo "LMS Go Backend - Available Commands"
	@echo "===================================="
	@echo "make build          - Build the application"
	@echo "make run            - Run the application"
	@echo "make test           - Run tests"
	@echo "make test-coverage  - Run tests with coverage report"
	@echo "make deps           - Download dependencies"
	@echo "make fmt            - Format code"
	@echo "make lint           - Run linter"
	@echo "make clean          - Clean build artifacts"
	@echo "make docker-up      - Start Docker containers"
	@echo "make docker-down    - Stop Docker containers"
	@echo "make db-migrate     - Run database migrations"
	@echo "make seed           - Seed database with test data"
	@echo "make dev            - Run in development mode with hot reload"

deps:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

lint:
	golangci-lint run ./...

build:
	CGO_ENABLED=0 go build -o lms-server .

run: build
	./lms-server

dev:
	go run main.go

test:
	go test -v -cover ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

seed:
	go run main.go seed.go seed

clean:
	rm -f lms-server coverage.out coverage.html

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

db-migrate:
	go run main.go migrate

docker-build:
	docker build -t lms-app:latest .

docker-run: docker-build
	docker run -p 8080:8080 --env-file .env lms-app:latest

.DEFAULT_GOAL := help

