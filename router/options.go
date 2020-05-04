package router

import (
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/etcd"
)

type Options struct {
	Namespace string
	mux       *runtime.ServeMuxDynamic
	Registry  registry.Registry
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		mux:      runtime.NewServeMuxDynamic(),
		Registry: etcd.NewRegistry(),
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithNamespace(ns string) Option {
	return func(o *Options) {
		o.Namespace = ns
	}
}

func WithMux(m *runtime.ServeMuxDynamic) Option {
	return func(o *Options) {
		o.mux = m
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}
