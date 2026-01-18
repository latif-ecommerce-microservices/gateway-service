.PHONY: run build test clean update-proto lint setup

-include .env
export $(shell sed 's/=.*//' .env)

run:
	@echo "Starting Gateway Service..."
	@go run cmd/server/main.go

update-proto:
	@echo "Updating user-service dependency..."
	go get github.com/latif-ecommerce-microservices/user-service@latest
	go mod tidy
	@echo "Done! User Service proto updated."

build:
	@echo "Building binary..."
	@go build -o bin/gateway cmd/server/main.go
	@echo "Build complete. Binary is in ./bin/gateway"

test:
	@go test -v ./...

clean:
	@rm -rf bin
	@echo "Cleanup done."

setup:
	@echo "Configuring GOPRIVATE for local development..."
	go env -w GOPRIVATE=github.com/latif-ecommerce-microservices/*
	@echo "Setup complete! Now you can run 'make update-proto' or 'make run'."