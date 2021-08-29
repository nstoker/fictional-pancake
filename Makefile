#!/usr/bin/make

include .env

MAIN=cmd/server/main.go

hello:
	@echo "Hello"

build: yarn-build
	@go build $(MAIN)

run: yarn-build
	@go run $(MAIN)

compile:
	@echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

yarn-build:
	@make -C frontend build
	
all: hello build

