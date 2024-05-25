
build:
	docker build -t registry.structx.local/structx/lightnode:v0.0.1 .

push:
	docker push registry.structx.local/structx/lightnode:v0.0.1

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...
