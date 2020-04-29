package main

import (
	"net/http"
	"time"

	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-chen/gateway/registry/etcd"
	"github.com/hb-chen/gateway/router"
	_ "github.com/hb-go/grpc-contrib/registry/micro"
	mregistry "github.com/micro/go-micro/v2/registry"
	metcd "github.com/micro/go-micro/v2/registry/etcd"
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
	grpcLogger, err := zapConf.Build(zap.AddCallerSkip(3))
	if err != nil {
		grpclog.Fatal(err)
	}
	grpc_zap.ReplaceGrpcLoggerV2(grpcLogger)

	//
	mregistry.DefaultRegistry = metcd.NewRegistry()

	runtime.DefaultContextTimeout = time.Second * 5
}

func main() {
	reg := etcd.NewRegistry()
	mux := runtime.NewServeMuxDynamic()

	r := router.NewRouter(router.WithMux(mux), router.WithRegistry(reg))
	defer r.Close()

	http.ListenAndServe(":8080", mux)
}
