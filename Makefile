
image:
	docker build -t structx/olivia:latest .

push:
	docker push trevatk/lightnode:latest

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...


rpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    proto/store/v1/local_store.proto