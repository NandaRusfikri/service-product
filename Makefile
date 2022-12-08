GO := @go
PROTOC := @protoc

protoc:
	${PROTOC} --proto_path=. --micro_out=. --go_out=. proto/*.proto
