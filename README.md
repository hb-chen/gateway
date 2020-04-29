# Gateway

## Example

**网关**
```shell script
$ cd cmd
$ go run gateway.go
```

**服务**
```shell script
$ cd example
$ go run service.go
```

gRPC 服务支持：
- Codec
    - `import _ "github.com/hb-chen/gateway/codec"`
- 注册中心
    - [github.com/hb-go/grpc-contrib/registry](https://github.com/hb-go/grpc-contrib/tree/master/registry)

- http://localhost:8080/v1/example/call
- http://localhost:8080/v1/example/call/hbchen
    
## Proto 工具

```shell script
$ git clone https://github.com/hb-chen/grpc-gateway.git $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
$ cd $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway
$ git checkout dynamic && git pull

$ go install ./protoc-gen-grpc-gateway 
```

    
```shell script
protoc --proto_path=.:$GOPATH/src \
--go_out=plugins=grpc:. \
--grpc-gateway_out=logtostderr=true,grpc_api_configuration=example/proto/gateway.yaml:. \
--hb-grpc_out=plugins=registry:. \
example/proto/service.proto
```

## Ref

- [micro/go-micro](https://github.com/micro/go-micro)
