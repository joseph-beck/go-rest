.PHONY: build run dev test

build:

run:

dev:
	@bash scripts/dev.sh

test: 
	@go test -cover ./...