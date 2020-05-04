package main

import (
	"fmt"
	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-chen/gateway/router"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"os"
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

	registry.DefaultRegistry = cache.New(etcd.NewRegistry())
}

func setup(ctx *cli.Context) error {
	provider := ctx.String("registry")
	// addr := ctx.String("registry_address")
	switch provider {
	case "etcd":
	default:
		return fmt.Errorf("registry provider:%v unsupported", provider)
	}
	return nil
}

func main() {
	flags := make([]cli.Flag, 0)
	flags = append(flags,
		&cli.StringFlag{
			Name:  "serve_addr",
			Value: ":8080",
			Usage: "serve address",
		},
		&cli.StringFlag{
			Name:  "service_version",
			Usage: "service version",
		},
		&cli.StringFlag{
			Name:  "registry",
			Value: "etcd",
			Usage: "micro registry provider, etcd",
		},
		&cli.StringFlag{
			Name:  "registry_address",
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
			serveAddr := ctx.String("serve_addr")
			// registryAddr := ctx.String("registry_address")

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
