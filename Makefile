.PHONY: grpcgen
grpcgen: ## generate protobuf files
	protoc --go_out=. --proto_path=./api/proto --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import \
	api/proto/*.proto