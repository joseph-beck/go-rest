.PHONY: build run dev test deploy

build:
	go build -o build/program/app cmd/httpd/main.go 

run:
	make build
	./build/program/app

dev:
	@bash scripts/dev.sh

test:
	go clean -testcache 
	go mod tidy
	go test -cover ./...

deploy:
	docker build --tag go-rest .
	docker run -p 8080:8080 --rm go-rest &