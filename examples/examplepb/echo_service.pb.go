// Code generated by protoc-gen-go.
// source: examples/examplepb/echo_service.proto
// DO NOT EDIT!

/*
Package examplepb is a generated protocol buffer package.

Echo Service

Echo Service API consists of a single service which returns
a message.

<!-- swagger extras start
{
  "info": {
    "title": "Echo Service",
    "version": "1.0",
    "contact": {
      "name": "gRPC-Gateway project",
      "url": "https://github.com/gengo/grpc-gateway",
      "email": "none@example.com"
    }
  },
  "host": "localhost",
  "externalDocs": {
    "url": "https://github.com/gengo/grpc-gateway",
    "description": "More about gRPC-Gateway"
  }
}
swagger extras end -->

It is generated from these files:
	examples/examplepb/echo_service.proto
	examples/examplepb/a_bit_of_everything.proto
	examples/examplepb/stream.proto
	examples/examplepb/flow_combination.proto

It has these top-level messages:
	SimpleMessage
	ABitOfEverything
	EmptyProto
	NonEmptyProto
	UnaryProto
	NestedProto
	SingleNestedProto
*/
package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SimpleMessage represents a simple message sent to the Echo service.
//
// <!-- swagger extras start
// {
//   "externalDocs": {
//     "url": "http://github.com/gengo/grpc-gateway",
//     "description": "Find out more about EchoService"
//   }
// }
// swagger extras end -->
type SimpleMessage struct {
	// Id represents the message identifier.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *SimpleMessage) Reset()                    { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string            { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()               {}
func (*SimpleMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SimpleMessage)(nil), "grpc.gateway.examples.examplepb.SimpleMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for EchoService service

type EchoServiceClient interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	//
	// <!-- swagger extras start
	// {
	//   "externalDocs": {
	//     "url": "http://github.com/gengo/grpc-gateway",
	//     "description": "Find out more about EchoService"
	//   }
	// }
	// swagger extras end -->
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.EchoService/EchoBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	// Echo method receives a simple message and returns it.
	//
	// The message posted as the id parameter will also be
	// returned.
	//
	// <!-- swagger extras start
	// {
	//   "externalDocs": {
	//     "url": "http://github.com/gengo/grpc-gateway",
	//     "description": "Find out more about EchoService"
	//   }
	// }
	// swagger extras end -->
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	// EchoBody method receives a simple message and returns it.
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("examples/examplepb/echo_service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x87, 0x32, 0x0a, 0x92, 0xf4, 0x53, 0x93, 0x33, 0xf2, 0xe3, 0x8b,
	0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0xd3, 0x8b,
	0x0a, 0x92, 0xf5, 0xd2, 0x13, 0x4b, 0x52, 0xcb, 0x13, 0x2b, 0xf5, 0x60, 0x7a, 0xf4, 0xe0, 0x7a,
	0xa4, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2,
	0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xda, 0x95, 0xe4, 0xb9, 0x78, 0x83, 0x33,
	0x41, 0x2a, 0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x80, 0x2c, 0xa3, 0x25, 0x4c, 0x5c, 0xdc, 0xae, 0x40, 0x6b,
	0x83, 0x21, 0xb6, 0x0a, 0xb5, 0x32, 0x72, 0xb1, 0x80, 0xf8, 0x42, 0x7a, 0x7a, 0x04, 0x6c, 0xd6,
	0x43, 0x31, 0x58, 0x8a, 0x44, 0xf5, 0x4a, 0xb2, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12, 0x57, 0x12,
	0xd5, 0x2f, 0x33, 0x84, 0x05, 0x01, 0x38, 0x00, 0xf4, 0xab, 0x33, 0x53, 0x6a, 0x85, 0x7a, 0x18,
	0xb9, 0x38, 0x40, 0xee, 0x70, 0xca, 0x4f, 0xa9, 0xa4, 0xb9, 0x5b, 0x14, 0xc0, 0x6e, 0x91, 0xc2,
	0x74, 0x4b, 0x7c, 0x12, 0xd0, 0x7a, 0x2b, 0x46, 0x2d, 0x27, 0xee, 0x28, 0x4e, 0xb8, 0xe6, 0x24,
	0x36, 0x70, 0xd8, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0x96, 0x37, 0xac, 0xc3, 0x01,
	0x00, 0x00,
}
