// Code generated by protoc-gen-hb-grpc. DO NOT EDIT.
// source: example/proto/service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	grpc "google.golang.org/grpc"
)

import (
	registry "github.com/hb-go/grpc-contrib/registry"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Export service desc

// Example desc
func ServiceDescExample() grpc.ServiceDesc {
	return _Example_serviceDesc
}

// gRPC registry service
// github.com/hb-go/grpc-contrib/registry

// Example registry service
var RegistryServiceExample = registry.Service{
	Name: _Example_serviceDesc.ServiceName,
	Methods: []*registry.Method{
		&registry.Method{
			Name: "Call",
		},
	},
}
