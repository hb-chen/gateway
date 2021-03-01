package router

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hb-go/grpc-contrib/registry"
)

type Options struct {
	namespace string
	mux       *runtime.ServeMuxDynamic
	registry  registry.Registry
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		mux:      runtime.NewServeMuxDynamic(),
		registry: &registry.MockRegistry{},
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

func WithNamespace(ns string) Option {
	return func(o *Options) {
		o.namespace = ns
	}
}

func WithMux(m *runtime.ServeMuxDynamic) Option {
	return func(o *Options) {
		o.mux = m
	}
}

func WithRegistry(r registry.Registry) Option {
	return func(o *Options) {
		o.registry = r
	}
}
