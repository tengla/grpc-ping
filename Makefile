
all:
	go build -o dist/client cmd/client/main.go
	go build -o dist/server cmd/server/main.go

protoc:
	protoc \
	-I=protos \
	--go_out=protos \
	--go_opt=paths=source_relative \
	--go-grpc_out=protos \
	--go-grpc_opt=paths=source_relative \
  protos/ping.proto
