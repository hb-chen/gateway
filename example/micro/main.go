package main

import (
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/hb-chen/gateway/codec"
	"github.com/micro/go-micro/v2/server/grpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/google/uuid"
	_ "github.com/hb-chen/gateway/codec"
	"github.com/hb-chen/gateway/example/micro/handler"
	gwRegistry "github.com/hb-chen/gateway/registry"
	gwEtcd "github.com/hb-chen/gateway/registry/etcd"
	"github.com/hb-go/grpc-contrib/registry"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
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
}

func main() {
	// TODO gw与micro proto生成代码冲突

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.example"),
		micro.Version("latest"),
		micro.Registry(memory.NewRegistry()), // 不需要 micro 的注册中心或者做一个 mock
	)

	// Initialise service
	addr := ""
	version := "v1"
	reg := gwEtcd.NewRegistry()
	gwService := example.RegistryServiceExample
	gwService.Version = version

	service.Init(
		micro.AfterStart(func() error {
			addr = service.Server().Options().Address

			// 服务注册
			if err := example.RegisterExample(registry.WithAddr(addr), registry.WithVersion(version)); err != nil {
				return err
			}

			// 网关注册
			gwService.Nodes = []*gwRegistry.Node{
				{
					Id:       gwService.Name + "-" + uuid.New().String(),
					Address:  addr,
					Metadata: nil,
				},
			}
			if err := reg.Register(&gwService); err != nil {
				return err
			}

			return nil
		}),
		micro.AfterStop(func() error {
			example.DeregisterExample(registry.WithAddr(addr), registry.WithVersion(version))
			if err := reg.Deregister(&gwService); err != nil {
				return err
			}

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
