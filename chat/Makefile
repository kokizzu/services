
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v5/cmd/protoc-gen-micro

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/chat.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/chat.proto

.PHONY: build
build:
	go build -o chat *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t chat:latest
