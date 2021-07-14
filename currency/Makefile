protos:
	protoc -I protos/ protos/currency.proto --go-grpc_out=./protos/currency
protos_full:
	protoc --go_out=./currency --go_opt=paths=source_relative \
	--go-grpc_out=./currency --go-grpc_opt=paths=source_relative \
	protos/currency.proto

.PHONY: protos, protos_full