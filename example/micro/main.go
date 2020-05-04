package main

import (
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/hb-chen/gateway/codec"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	"github.com/micro/go-micro/v2/server/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/google/uuid"
	_ "github.com/hb-chen/gateway/codec"
	"github.com/hb-chen/gateway/example/micro/handler"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry/memory"
	"github.com/micro/go-micro/v2/util/log"
	"google.golang.org/grpc/grpclog"

	example "github.com/hb-chen/gateway/example/micro/proto/example"
)

func init() {
	zapEncoderConf := zap.NewDevelopmentEncoderConfig()
	zapEncoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConf := zap.NewDevelopmentConfig()
	zapConf.EncoderConfig = zapEncoderConf
	zapConf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	grpcLogger, err := zapConf.Build(zap.AddCallerSkip(3))
	if err != nil {
		grpclog.Fatal(err)
	}
	grpcZap.ReplaceGrpcLoggerV2(grpcLogger)

	registry.DefaultRegistry = cache.New(etcd.NewRegistry())
}

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.example"),
		micro.Version("latest"),
		micro.Registry(memory.NewRegistry()), // 不需要 micro 的注册中心或者做一个 mock
	)

	// Initialise service
	addr := ""
	version := "v1"
	exampleService := example.GatewayServiceExample
	exampleService.Version = version

	service.Init(
		micro.AfterStart(func() error {
			addr = service.Server().Options().Address

			// 服务注册
			exampleService.Nodes = []*registry.Node{
				{
					Id:       exampleService.Name + "-" + uuid.New().String(),
					Address:  addr,
					Metadata: nil,
				},
			}
			if err := registry.Register(&exampleService); err != nil {
				return err
			}

			return nil
		}),
		micro.AfterStop(func() error {
			// 服务注销
			registry.Deregister(&exampleService)

			return nil
		}),
	)

	// 网关 Codec
	jsonCodec := codec.JsonCodec{}
	service.Server().Init(
		grpc.Codec("application/grpc+"+jsonCodec.Name(), jsonCodec),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
