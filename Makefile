GO := @go
PROTOC := @protoc

protoc:
	${PROTOC} --proto_path=. --micro_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_out=. proto/product/*.proto
