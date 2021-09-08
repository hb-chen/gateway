package router

import (
	"reflect"
	"sync"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hb-go/grpc-contrib/registry"
	"github.com/hb-go/grpc-contrib/registry/cache"
)

func TestNewOptions(t *testing.T) {
	type args struct {
		opts []Option
	}

	want := NewOptions()
	want.namespace = "ns"
	want.mux = &runtime.ServeMuxDynamic{}
	want.registry = &registry.MockRegistry{}

	tests := []struct {
		name string
		args args
		want Options
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{
				opts: []Option{
					WithNamespace(want.namespace),
					WithMux(want.mux),
					WithRegistry(want.registry),
				},
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRouter(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *registryRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithMux(t *testing.T) {
	type args struct {
		m *runtime.ServeMuxDynamic
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithMux(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithMux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithNamespace(t *testing.T) {
	type args struct {
		ns string
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithNamespace(tt.args.ns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRegistry(t *testing.T) {
	type args struct {
		r registry.Registry
	}
	tests := []struct {
		name string
		args args
		want Option
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRegistry(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRegistry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRouter(t *testing.T) {
	type args struct {
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *registryRouter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRouter(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registryRouter_Close(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}
			if err := r.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_registryRouter_Options(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	tests := []struct {
		name   string
		fields fields
		want   Options
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}
			if got := r.Options(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Options() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registryRouter_handler(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	type args struct {
		serviceName string
		method      string
		versions    []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   runtime.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}
			if got := r.handler(tt.args.serviceName, tt.args.method, tt.args.versions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registryRouter_isClosed(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}
			if got := r.isClosed(); got != tt.want {
				t.Errorf("isClosed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_registryRouter_process(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	type args struct {
		res *registry.Result
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}
			_ = r
		})
	}
}

func Test_registryRouter_refresh(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}

			_ = r
		})
	}
}

func Test_registryRouter_store(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	type args struct {
		services []*registry.Service
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}

			_ = r
		})
	}
}

func Test_registryRouter_watch(t *testing.T) {
	type fields struct {
		exit    chan bool
		opts    Options
		rc      cache.Cache
		RWMutex sync.RWMutex
		eps     map[string]*registry.Service
		routes  map[string]*Route
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &registryRouter{
				exit:    tt.fields.exit,
				opts:    tt.fields.opts,
				rc:      tt.fields.rc,
				RWMutex: tt.fields.RWMutex,
				eps:     tt.fields.eps,
				routes:  tt.fields.routes,
			}

			_ = r
		})
	}
}
