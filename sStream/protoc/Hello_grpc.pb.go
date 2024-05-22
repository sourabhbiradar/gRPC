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
	SStream_ServerReply_FullMethodName = "/SStream/ServerReply"
)

// SStreamClient is the client API for SStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SStreamClient interface {
	ServerReply(ctx context.Context, in *CReqt, opts ...grpc.CallOption) (SStream_ServerReplyClient, error)
}

type sStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewSStreamClient(cc grpc.ClientConnInterface) SStreamClient {
	return &sStreamClient{cc}
}

func (c *sStreamClient) ServerReply(ctx context.Context, in *CReqt, opts ...grpc.CallOption) (SStream_ServerReplyClient, error) {
	stream, err := c.cc.NewStream(ctx, &SStream_ServiceDesc.Streams[0], SStream_ServerReply_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &sStreamServerReplyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SStream_ServerReplyClient interface {
	Recv() (*SResps, error)
	grpc.ClientStream
}

type sStreamServerReplyClient struct {
	grpc.ClientStream
}

func (x *sStreamServerReplyClient) Recv() (*SResps, error) {
	m := new(SResps)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SStreamServer is the server API for SStream service.
// All implementations must embed UnimplementedSStreamServer
// for forward compatibility
type SStreamServer interface {
	ServerReply(*CReqt, SStream_ServerReplyServer) error
	mustEmbedUnimplementedSStreamServer()
}

// UnimplementedSStreamServer must be embedded to have forward compatible implementations.
type UnimplementedSStreamServer struct {
}

func (UnimplementedSStreamServer) ServerReply(*CReqt, SStream_ServerReplyServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerReply not implemented")
}
func (UnimplementedSStreamServer) mustEmbedUnimplementedSStreamServer() {}

// UnsafeSStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SStreamServer will
// result in compilation errors.
type UnsafeSStreamServer interface {
	mustEmbedUnimplementedSStreamServer()
}

func RegisterSStreamServer(s grpc.ServiceRegistrar, srv SStreamServer) {
	s.RegisterService(&SStream_ServiceDesc, srv)
}

func _SStream_ServerReply_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CReqt)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SStreamServer).ServerReply(m, &sStreamServerReplyServer{stream})
}

type SStream_ServerReplyServer interface {
	Send(*SResps) error
	grpc.ServerStream
}

type sStreamServerReplyServer struct {
	grpc.ServerStream
}

func (x *sStreamServerReplyServer) Send(m *SResps) error {
	return x.ServerStream.SendMsg(m)
}

// SStream_ServiceDesc is the grpc.ServiceDesc for SStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SStream",
	HandlerType: (*SStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerReply",
			Handler:       _SStream_ServerReply_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "Hello.proto",
}
