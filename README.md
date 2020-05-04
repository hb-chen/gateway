# Gateway

## Example

**网关**
```shell script
$ cd cmd
$ go run main.go
```

**gRPC 服务**
```shell script
$ cd example
$ go run service.go
```

**go-micro 的 gRPC 服务**
```shell script
$ cd example/micro
$ go run main.go
```

gRPC 服务支持：
- Codec
    - `import _ "github.com/hb-chen/gateway/codec"`
    - go-micro 使用 gRPC server 的 Codec() option
        `grpc.Codec("application/grpc+"+jsonCodec.Name(), jsonCodec),`
- 注册中心
    - [github.com/hb-go/grpc-contrib/registry](https://github.com/hb-go/grpc-contrib/tree/master/registry)

**测试接口**
- http://localhost:8080/v1/example/call
- http://localhost:8080/v1/example/call/hbchen
    
## Proto 工具

[protoc-gen-hb-grpc](https://github.com/hb-go/grpc-contrib)
[protoc-gen-hb-grpc-gateway](https://github.com/hb-chen/gateway)

```shell script
$ go install github.com/hb-go/grpc-contrib/protoc-gen-hb-grpc
$ go install github.com/hb-chen/gateway/protoc-gen-hb-grpc-gateway 
```
    
```shell script
protoc --proto_path=.:$GOPATH/src \
--go_out=plugins=grpc:. \
--hb-grpc-gateway_out=logtostderr=true,grpc_api_configuration=example/proto/gateway.yaml:. \
--hb-grpc_out=plugins=desc+registry:. \
example/proto/service.proto
```

## Ref

- [micro/go-micro](https://github.com/micro/go-micro)
