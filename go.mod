module github.com/hb-chen/gateway

go 1.13

require (
	github.com/bitly/go-simplejson v0.5.0
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.14.4
	github.com/hb-go/grpc-contrib v0.0.0-20200429050913-3e94c480880b
	github.com/micro/go-micro/v2 v2.4.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/urfave/cli/v2 v2.2.0 // indirect
	go.uber.org/zap v1.15.0
	google.golang.org/grpc v1.26.0
)

replace github.com/grpc-ecosystem/grpc-gateway v1.14.4 => github.com/hb-chen/grpc-gateway v1.14.4-dynamic
