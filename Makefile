.PHONY: build run dev test

build:

run:
	@bash scripts/run.sh

dev:
	@bash scripts/dev.sh

test: 
	@go test -cover ./...