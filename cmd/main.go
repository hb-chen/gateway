package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	grpcZap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/cache"
	"github.com/hb-go/grpc-contrib/registry/consul"
	"github.com/hb-go/grpc-contrib/registry/etcd"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/hb-chen/gateway/v2/pkg/router"
)

const (
	logCallerSkip = 2
)

func initLogger(path, level string, debug, e bool) error {
	logLevel := zapcore.WarnLevel
	err := logLevel.UnmarshalText([]byte(level))
	if err != nil {
		return err
	}

	writer := logWriter(path)
	if e {
		stderr, close, err := zap.Open("stderr")
		if err != nil {
			close()
			return err
		}
		writer = stderr
	}

	encoder := logEncoder(debug)
	core := zapcore.NewCore(encoder, writer, logLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(logCallerSkip))
	grpcZap.ReplaceGrpcLoggerV2(logger)

	return nil
}

func logEncoder(debug bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	if debug {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	}

	return zapcore.NewConsoleEncoder(encoderConfig)
}

func logWriter(path string) zapcore.WriteSyncer {
	path = strings.TrimRight(path, "/")
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path + "/gateway.log",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     7,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
func setup(ctx *cli.Context) error {
	// 日志
	debug := ctx.Bool("debug")
	e := ctx.Bool("e")
	level := ctx.String("log_level")
	path := ctx.String("log_path")
	if err := initLogger(path, level, debug, e); err != nil {
		return err
	}

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
			Value: "v2.3.0",
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
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "log for debug.",
		},
		&cli.BoolFlag{
			Name:  "e",
			Usage: "log to stderr.",
		},
		&cli.StringFlag{
			Name:  "log_level",
			Usage: "log level, debug info warn or error",
			Value: "warn",
		},
		&cli.StringFlag{
			Name:  "log_path",
			Usage: "log path.",
			Value: "./log",
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
				_, _ = w.Write([]byte(`{"name":"` + name + `","version": "` + version + `"}`))
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
