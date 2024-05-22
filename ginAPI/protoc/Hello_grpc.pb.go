// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: Hello.proto

package __

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
	WithAPI_ServerResp_FullMethodName = "/WithAPI/ServerResp"
)

// WithAPIClient is the client API for WithAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WithAPIClient interface {
	ServerResp(ctx context.Context, in *Creqt, opts ...grpc.CallOption) (*Sresp, error)
}

type withAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewWithAPIClient(cc grpc.ClientConnInterface) WithAPIClient {
	return &withAPIClient{cc}
}

func (c *withAPIClient) ServerResp(ctx context.Context, in *Creqt, opts ...grpc.CallOption) (*Sresp, error) {
	out := new(Sresp)
	err := c.cc.Invoke(ctx, WithAPI_ServerResp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WithAPIServer is the server API for WithAPI service.
// All implementations must embed UnimplementedWithAPIServer
// for forward compatibility
type WithAPIServer interface {
	ServerResp(context.Context, *Creqt) (*Sresp, error)
	mustEmbedUnimplementedWithAPIServer()
}

// UnimplementedWithAPIServer must be embedded to have forward compatible implementations.
type UnimplementedWithAPIServer struct {
}

func (UnimplementedWithAPIServer) ServerResp(context.Context, *Creqt) (*Sresp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerResp not implemented")
}
func (UnimplementedWithAPIServer) mustEmbedUnimplementedWithAPIServer() {}

// UnsafeWithAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WithAPIServer will
// result in compilation errors.
type UnsafeWithAPIServer interface {
	mustEmbedUnimplementedWithAPIServer()
}

func RegisterWithAPIServer(s grpc.ServiceRegistrar, srv WithAPIServer) {
	s.RegisterService(&WithAPI_ServiceDesc, srv)
}

func _WithAPI_ServerResp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Creqt)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WithAPIServer).ServerResp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WithAPI_ServerResp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WithAPIServer).ServerResp(ctx, req.(*Creqt))
	}
	return interceptor(ctx, in, info, handler)
}

// WithAPI_ServiceDesc is the grpc.ServiceDesc for WithAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WithAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WithAPI",
	HandlerType: (*WithAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerResp",
			Handler:    _WithAPI_ServerResp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Hello.proto",
}