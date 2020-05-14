package main

import (
	"fmt"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-chen/gateway/router"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"github.com/hb-go/grpc-contrib/registry/consul"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"os"
	"strings"
)

func init() {
	zapEncoderConf := zap.NewDevelopmentEncoderConfig()
	zapEncoderConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
	zapConf := zap.NewDevelopmentConfig()
	zapConf.EncoderConfig = zapEncoderConf
	zapConf.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, err := zapConf.Build(zap.AddCallerSkip(3))
	if err != nil {
		grpclog.Fatal(err)
	}
	grpcZap.ReplaceGrpcLoggerV2(logger)
}

func setup(ctx *cli.Context) error {
	provider := ctx.String("grpc_registry")
	addr := ctx.String("grpc_registry_address")
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
			Name:  "server_address",
			Value: ":8080",
			Usage: "server address",
		},
		&cli.StringFlag{
			Name:  "service_version",
			Usage: "service version",
		},
		&cli.StringFlag{
			Name:  "grpc_registry",
			Value: "etcd",
			Usage: "micro registry provider, etcd",
		},
		&cli.StringFlag{
			Name:  "grpc_registry_address",
			Usage: "micro registry address",
		},
	)

	app := &cli.App{
		Name:        "gateway",
		Description: "gRPC gateway",
		Flags:       flags,
		Before: func(ctx *cli.Context) error {
			return setup(ctx)
		},
		Action: func(ctx *cli.Context) error {
			serveAddr := ctx.String("server_address")

			mux := runtime.NewServeMuxDynamic()
			r := router.NewRouter(router.WithMux(mux), router.WithRegistry(registry.DefaultRegistry))
			defer r.Close()

			return http.ListenAndServe(serveAddr, mux)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		grpclog.Fatal(err)
	}
}
