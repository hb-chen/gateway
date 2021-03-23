GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc --proto_path=.:$GOPATH/src \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=logtostderr=true,grpc_api_configuration=example/proto/gateway.yaml:. \
    --hb-grpc-gateway_out=logtostderr=true,grpc_api_configuration=example/proto/gateway.yaml:. \
    --hb-grpc_out=plugins=desc+registry:. \
    example/proto/service.proto

.PHONY: test
test:
	go test -race -cover -v ./...

.PHONY: run
run:
	go run cmd/main.go --grpc_registry=etcd --log_level=info --debug -e

.PHONY: build
build:
	CGO_ENABLED=0 GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w' -o ./bin/gateway cmd/main.go

.PHONY: docker
docker: build
	docker build . -t $(tag)
