# Gateway

## Example

```shell script
$ cd example

# 网关
$ go run gateway.go

# 服务
$ go run service.go
```

http://localhost:8080/v1/example/call
http://localhost:8080/v1/example/call/hbchen
    
## Proto 工具

```shell script
$ git clone https://github.com/hb-chen/grpc-gateway.git $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway

$ git checkout dynamic && git pull

$ go install ./protoc-gen-grpc-gateway 
```

    
```shell script
protoc --proto_path=.:$GOPATH/src \
--go_out=plugins=grpc:. \
--grpc-gateway_out=logtostderr=true,grpc_api_configuration=proto/gateway.yaml:. \
--hb-grpc_out=plugins=registry:. \
proto/service.proto
```

## Ref

- [micro/go-micro](https://github.com/micro/go-micro)
