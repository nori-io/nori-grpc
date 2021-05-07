.PHONY: grpcgen
grpcgen: ## generate protobuf files
	protoc --go_out=plugins=grpc:./ -I api/proto api/proto/*.proto