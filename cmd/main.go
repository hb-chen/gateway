package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

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
	case "mock":
		grpclog.Warningf("registry mock")
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
			Name:  "name",
			Value: "gRPC gateway",
			Usage: "gateway name",
		},
		&cli.StringFlag{
			Name:  "version",
			Value: "v1.0.0",
			Usage: "gateway version",
		},
		&cli.StringFlag{
			Name:  "address",
			Value: ":8080",
			Usage: "serve address",
		},
		&cli.StringFlag{
			Name:  "grpc_registry",
			Value: "mock",
			Usage: "registry provider, etcd or consul",
		},
		&cli.StringFlag{
			Name:  "grpc_registry_address",
			Usage: "registry address",
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
			name := ctx.String("name")
			version := ctx.String("version")
			addr := ctx.String("address")

			mux := runtime.NewServeMuxDynamic()

			pat, _ := runtime.NewPattern(1, []int{1, 0}, []string{""}, "")
			mux.Handle("GET", pat, func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"` + name + `","version": "` + version + `"}`))
			})

			r := router.NewRouter(router.WithMux(mux), router.WithRegistry(registry.DefaultRegistry))
			defer r.Close()

			return http.ListenAndServe(addr, mux)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		grpclog.Fatal(err)
	}
}
