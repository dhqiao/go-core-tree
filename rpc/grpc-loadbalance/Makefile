all:
	@echo todo

build-protos:
	protoc -I./proto --go_out=plugins=grpc:pb $(wildcard ./proto/*.proto)