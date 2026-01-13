GO_WORKSPACE := .

build_proto:
	protoc --proto_path=proto proto/*.proto --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "Protoc compile done"