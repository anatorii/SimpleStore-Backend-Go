# Переменные
BINARY_NAME=shop
MAIN_PATH=cmd/shop/main.go
SWAGGER_DOCS=docs

# Go параметры
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Пути
GOPATH=$(shell go env GOPATH)
SWAG=$(GOPATH)/bin/swag

.PHONY: all build server clean swagger deps

swagger:
	$(SWAG) init -g cmd/shop/main.go -o docs --parseDependency --parseInternal

server: swagger
	go build -o shop cmd/shop/main.go
	./bin/shop

build: swagger
	go build -o shop cmd/shop/main.go

help:
	@echo "Available commands:"
	@echo "  make deps     - Install dependencies"
	@echo "  make swagger  - Generate Swagger documentation"
	@echo "  make server   - Run the application"
	@echo "  make build    - Build the application"	
	@echo "  make clean    - Clean the application"	

deps:
	go mod download
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

clean:
	rm -rf bin/
	rm -rf $(SWAGGER_DOCS)
	$(GOCMD) clean
