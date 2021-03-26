package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/google/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"github.com/hb-go/grpc-contrib/registry/consul"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"github.com/hb-chen/gateway/v2/example/proto"
	_ "github.com/hb-chen/gateway/v2/pkg/codec"
	"github.com/hb-chen/gateway/v2/pkg/util"
	mNet "github.com/hb-chen/gateway/v2/pkg/util/net"
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

type exampleService struct {
	proto.UnimplementedExampleServer
}

func (*exampleService) Call(_ context.Context, req *proto.Request) (*proto.Response, error) {
	grpclog.Infof("example service request call")
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, `req.Name=""`)
	}
	return &proto.Response{Code: 0, Msg: "Hello " + req.Name}, nil
}

func setup(ctx *cli.Context) error {
	provider := ctx.String("grpc_registry")
	addr := ctx.String("grpc_registry_address")
	grpclog.Infof("registry provider: %v", provider)
	switch provider {
	case "etcd":
		registry.DefaultRegistry = cache.New(etcd.NewRegistry(registry.Addrs(strings.Split(addr, ",")...)))
	case "consul":
		registry.DefaultRegistry = cache.New(consul.NewRegistry(registry.Addrs(strings.Split(addr, ",")...)))
	default:
		return fmt.Errorf("registry provider:%v unsupported", provider)
	}
	return nil
}

func main() {
	flags := make([]cli.Flag, 0)
	flags = append(flags,
		&cli.StringFlag{
			Name:  "grpc_registry",
			Value: "etcd",
			Usage: "registry provider, etcd or consul",
		},
		&cli.StringFlag{
			Name:  "grpc_registry_address",
			Usage: "registry address",
		},
	)

	// 服务
	version := "v1"
	service := proto.GatewayServiceExample
	service.Version = version

	s := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	app := &cli.App{
		Name:        "example",
		Description: "example service",
		Flags:       flags,
		Before: func(ctx *cli.Context) error {
			grpclog.Infof("before func")
			return setup(ctx)
		},
		Action: func(ctx *cli.Context) error {
			l, err := mNet.Listen(":", func(addr string) (net.Listener, error) {
				return net.Listen("tcp", addr)
			})
			if err != nil {
				grpclog.Fatalf("failed to listen: %v", err)
			}

			address, err := util.Address(l.Addr().String())
			if err != nil {
				grpclog.Fatalf("address error: %v", err)
			}

			// 服务注册
			service.Nodes = []*registry.Node{
				{
					Id:       service.Name + "-" + uuid.New().String(),
					Address:  address,
					Metadata: nil,
				},
			}
			if err := registry.Register(&service); err != nil {
				grpclog.Fatal(err)
			}

			proto.RegisterExampleServer(s, &exampleService{})

			// Register reflection service on gRPC server.
			reflection.Register(s)

			grpclog.Infof("grpc serve addr: %v", l.Addr().String())
			return s.Serve(l)
		},
	}

	c := make(chan os.Signal, 2)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		grpclog.Infof("exit signal")

		err := registry.Deregister(&service)
		if err != nil {
			grpclog.Errorf("service deregister failed: %v", err)
		}

		s.GracefulStop()

		grpclog.Infof("exit")
		time.Sleep(3 * time.Second) // custom behavior on the user side
		os.Exit(0)
	}()

	err := app.Run(os.Args)
	if err != nil {
		grpclog.Fatal(err)
	}
}
