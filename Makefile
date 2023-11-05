.PHONY: build
build:
	go build -o ./bin -v ./cmd/apiserver

start:
	./bin/apiserver.exe

dcb:
	docker build ./ --tag social-network-ws-api

.DEFAULT_GOAL := build