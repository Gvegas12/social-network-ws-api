.PHONY: build
build:
	go build -o ./bin -v ./cmd/apiserver

start:
	./bin/apiserver.exe

.DEFAULT_GOAL := build