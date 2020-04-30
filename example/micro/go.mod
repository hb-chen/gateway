module github.com/hb-chen/gateway/example/micro

go 1.13

require (
	github.com/golang/protobuf v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/hb-chen/gateway v0.0.0-20200429072353-1e55f50970c2
	github.com/hb-go/grpc-contrib v0.0.0-20200429050913-3e94c480880b
	github.com/micro/go-micro/v2 v2.4.0
	google.golang.org/grpc v1.26.0
)

replace github.com/hb-chen/gateway v0.0.0-20200429072353-1e55f50970c2 => ../../../gateway
