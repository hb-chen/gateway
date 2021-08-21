module github.com/hb-chen/gateway/v2

go 1.15

replace github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0 => github.com/hb-chen/grpc-gateway/v2 v2.3.0-dynamic

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/grpc-ecosystem/grpc-gateway v1.14.6
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.3.0
	github.com/hb-go/grpc-contrib v0.0.2-0.20210301130204-972628bec1d8
	github.com/urfave/cli/v2 v2.2.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.1-0.20201208041424-160c7477e0e8
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
