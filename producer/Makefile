run:
	go run .

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/stats_service.proto --experimental_allow_proto3_optional

.PHONY: proto
