package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	_ "github.com/hb-chen/gateway/codec"
	"github.com/hb-chen/gateway/example/proto"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	mNet "github.com/micro/go-micro/v2/util/net"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

type exampleService struct {
}

func (*exampleService) Call(_ context.Context, req *proto.Request) (*proto.Response, error) {
	grpclog.Infof("example service request call")
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, `req.Name=""`)
	}
	return &proto.Response{Code: 0, Msg: "Hello " + req.Name}, nil
}

func main() {
	l, err := mNet.Listen(":", func(addr string) (net.Listener, error) {
		return net.Listen("tcp", addr)
	})
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	version := "v1"

	// 网关注册
	service := proto.GatewayServiceExampleService
	service.Version = version
	service.Nodes = []*registry.Node{
		{
			Id:       service.Name + "-" + uuid.New().String(),
			Address:  l.Addr().String(),
			Metadata: nil,
		},
	}

	if err := registry.Register(&service); err != nil {
		grpclog.Fatal(err)
	}

	s := grpc.NewServer(
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	proto.RegisterExampleServiceServer(s, &exampleService{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGKILL, syscall.SIGINT)

	exitCh := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		select {
		case <-exitCh:
			grpclog.Infof("server exit")
		case sig := <-ch:
			grpclog.Infof("receive signal: %v", sig.String())
			s.GracefulStop()
		}

		defer registry.Deregister(&service)

		wg.Done()
	}(wg)

	grpclog.Infof("grpc serve addr: %v", l.Addr().String())
	if err := s.Serve(l); err != nil {
		grpclog.Fatalf("grpc failed to serve: %v", err)
	} else {
		grpclog.Infof("grpc serve end")
	}

	close(exitCh)
	wg.Wait()
}
