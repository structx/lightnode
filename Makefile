
build:
	docker build -t trevatk/lightnode:latest .

push:
	docker push trevatk/lightnode:latest

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...


rpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/raft/v1/raft_service.proto