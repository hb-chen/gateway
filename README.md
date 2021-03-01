# Gateway

## Helm install

```
helm upgrade --install grpc-gateway ./helm --namespace grpc-gateway --no-hooks
```

## Example

**网关**
```shell script
$ cd cmd
$ go run main.go --grpc_registry=etcd --server_address=:8080
```

**gRPC 服务**
```shell script
$ cd example/grpc
$ go run main.go --grpc_registry=etcd
```

**~~go-micro 的 gRPC 服务~~**
```shell script
$ cd example/micro
$ go run main.go --grpc_registry=etcd
```

gRPC 服务支持：
- Codec
    - `import _ "github.com/hb-chen/gateway/v2/codec"`
    - go-micro 使用 gRPC server 的 Codec() option
        `grpc.Codec("application/grpc+"+jsonCodec.Name(), jsonCodec),`
- 注册中心
    - [github.com/hb-go/grpc-contrib/registry](https://github.com/hb-go/grpc-contrib/tree/master/registry)

**测试接口**

```shell script
# POST
curl -X POST -d '{"name":"hbchen"}' http://localhost:8080/v1/example/call
{"code":"0","msg":"Hello hbchen"}

# GET
curl http://localhost:8080/v1/example/call/hbchen
{"code":"0","msg":"Hello hbchen"}
```
    
## Proto 工具

[protoc-gen-hb-grpc](https://github.com/hb-go/grpc-contrib)
[protoc-gen-hb-grpc-gateway](https://github.com/hb-chen/gateway)

```shell script
$ go install github.com/hb-go/grpc-contrib/protoc-gen-hb-grpc
$ go install github.com/hb-chen/gateway/v2/protoc-gen-hb-grpc-gateway 
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
