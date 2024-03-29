// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: userrpc.proto

package userrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StreamGreeter_SayHello_FullMethodName = "/stream.StreamGreeter/SayHello"
)

// StreamGreeterClient is the client API for StreamGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamGreeterClient interface {
	SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error)
}

type streamGreeterClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamGreeterClient(cc grpc.ClientConnInterface) StreamGreeterClient {
	return &streamGreeterClient{cc}
}

func (c *streamGreeterClient) SayHello(ctx context.Context, in *SayHelloReq, opts ...grpc.CallOption) (*SayHelloResp, error) {
	out := new(SayHelloResp)
	err := c.cc.Invoke(ctx, StreamGreeter_SayHello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamGreeterServer is the server API for StreamGreeter service.
// All implementations must embed UnimplementedStreamGreeterServer
// for forward compatibility
type StreamGreeterServer interface {
	SayHello(context.Context, *SayHelloReq) (*SayHelloResp, error)
	mustEmbedUnimplementedStreamGreeterServer()
}

// UnimplementedStreamGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedStreamGreeterServer struct {
}

func (UnimplementedStreamGreeterServer) SayHello(context.Context, *SayHelloReq) (*SayHelloResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedStreamGreeterServer) mustEmbedUnimplementedStreamGreeterServer() {}

// UnsafeStreamGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamGreeterServer will
// result in compilation errors.
type UnsafeStreamGreeterServer interface {
	mustEmbedUnimplementedStreamGreeterServer()
}

func RegisterStreamGreeterServer(s grpc.ServiceRegistrar, srv StreamGreeterServer) {
	s.RegisterService(&StreamGreeter_ServiceDesc, srv)
}

func _StreamGreeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamGreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamGreeter_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamGreeterServer).SayHello(ctx, req.(*SayHelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamGreeter_ServiceDesc is the grpc.ServiceDesc for StreamGreeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamGreeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.StreamGreeter",
	HandlerType: (*StreamGreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _StreamGreeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userrpc.proto",
}
