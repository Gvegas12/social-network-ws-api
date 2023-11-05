.PHONY: build
build:
	go build -o ./bin -v ./cmd/apiserver

start:
	./bin/apiserver.exe

dev:
	go run ./cmd/apiserver

dcb:
	docker build ./ --tag social-network-ws-api

.DEFAULT_GOAL := build