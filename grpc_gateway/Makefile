.PHONY: protos


proto_hello:
	protoc -I proto/ hello.proto --go_out=plugins=grpc:proto --grpc-gateway_out=:proto --openapiv2_out=:swagger