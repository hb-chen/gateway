
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=.:$$GOPATH/src \
	--micro_out=. \
    --go_out=plugins=grpc:. \
    --hb-grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/example/gateway.yaml:. \
    --hb-grpc_out=plugins=registry:. \
    proto/example/example.proto

.PHONY: build
build: proto

	go build -o example-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t example-srv:latest
