package router

import (
	"reflect"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hb-go/grpc-contrib/registry"
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
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOptions(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOptions() = %v, want %v", got, tt.want)
			}
		})
	}
}
