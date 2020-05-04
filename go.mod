module github.com/hb-chen/gateway

go 1.13

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/hb-go/grpc-contrib v0.0.0-20200504062014-d5f5132289c5
	github.com/micro/go-micro/v2 v2.4.0
	github.com/urfave/cli/v2 v2.2.0
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.26.0
)

replace github.com/grpc-ecosystem/grpc-gateway v1.14.4 => github.com/hb-chen/grpc-gateway v1.14.4-dynamic
