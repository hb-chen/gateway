// Code generated by protoc-gen-hb-grpc. DO NOT EDIT.
// source: proto/service.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// gRPC Registry
// github.com/hb-go/grpc-contrib/registry

// ExampleService registry
func TargetExampleService(opts ...registry.Option) string {
	return registry.NewTarget(&_ExampleService_serviceDesc, opts...)
}

func RegisterExampleService(opts ...registry.Option) error {
	return registry.Register(&_ExampleService_serviceDesc, opts...)
}

func DeregisterExampleService(opts ...registry.Option) {
	registry.Deregister(&_ExampleService_serviceDesc, opts...)
}
