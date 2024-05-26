
build:
	docker build -t registry.structx.local/structx/lightnode:v0.0.1 .

push:
	docker push registry.structx.local/structx/lightnode:v0.0.1

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...


rpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/raft/v1/raft_service.proto