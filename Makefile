
build:
	docker build -t structx/lightnode:latest .

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...

rpc:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/k2/v1/k2_service.proto