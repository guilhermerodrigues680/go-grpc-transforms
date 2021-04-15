.PHONY: default
default: cross

.PHONY: clearbin
clearbin:
	rm -rf ./bin

build: cmd clearbin
	GOOS=linux GOARCH=amd64 go build -v -o bin/main-server-linux-amd64 ./cmd/server
	GOOS=linux GOARCH=amd64 go build -v -o bin/main-client-linux-amd64 ./cmd/client

.PHONY: cross
cross: cmd clearbin
	go build -v -o bin/main-server ./cmd/server
	go build -v -o bin/main-client ./cmd/client

.PHONY: protobuffer
protobuffer:
	protoc --go_out=./transport/grpc --go-grpc_out=./transport/grpc ./protobuffer-files/*.proto

.PHONY: protobuffer-with-transform
protobuffer-with-transform:
	protoc --gotemplate_out=template_dir=./gen/template:./gen ./protobuffer-files/*.proto
