.PHONY: pb

pb: 
	protoc --go_out=plugins=grpc:. *.proto