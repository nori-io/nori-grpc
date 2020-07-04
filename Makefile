# Nori gRPC Makefile

grpc-gen: ## generate protobuf
	@protoc --proto_path=api/grpc -I=api/grpc --go_out=plugins=grpc:. api/grpc/*.proto
.PHONY: protoc-generate