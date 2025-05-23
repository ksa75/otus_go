// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package main

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ElectionsClient is the client API for Elections service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectionsClient interface {
	SubmitVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*empty.Empty, error)
	Internal(ctx context.Context, opts ...grpc.CallOption) (Elections_InternalClient, error)
}

type electionsClient struct {
	cc grpc.ClientConnInterface
}

func NewElectionsClient(cc grpc.ClientConnInterface) ElectionsClient {
	return &electionsClient{cc}
}

func (c *electionsClient) SubmitVote(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/elections_with_admin.Elections/SubmitVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electionsClient) Internal(ctx context.Context, opts ...grpc.CallOption) (Elections_InternalClient, error) {
	stream, err := c.cc.NewStream(ctx, &Elections_ServiceDesc.Streams[0], "/elections_with_admin.Elections/Internal", opts...)
	if err != nil {
		return nil, err
	}
	x := &electionsInternalClient{stream}
	return x, nil
}

type Elections_InternalClient interface {
	Send(*Vote) error
	Recv() (*StatsVote, error)
	grpc.ClientStream
}

type electionsInternalClient struct {
	grpc.ClientStream
}

func (x *electionsInternalClient) Send(m *Vote) error {
	return x.ClientStream.SendMsg(m)
}

func (x *electionsInternalClient) Recv() (*StatsVote, error) {
	m := new(StatsVote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ElectionsServer is the server API for Elections service.
// All implementations must embed UnimplementedElectionsServer
// for forward compatibility
type ElectionsServer interface {
	SubmitVote(context.Context, *Vote) (*empty.Empty, error)
	Internal(Elections_InternalServer) error
	mustEmbedUnimplementedElectionsServer()
}

// UnimplementedElectionsServer must be embedded to have forward compatible implementations.
type UnimplementedElectionsServer struct {
}

func (UnimplementedElectionsServer) SubmitVote(context.Context, *Vote) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitVote not implemented")
}
func (UnimplementedElectionsServer) Internal(Elections_InternalServer) error {
	return status.Errorf(codes.Unimplemented, "method Internal not implemented")
}
func (UnimplementedElectionsServer) mustEmbedUnimplementedElectionsServer() {}

// UnsafeElectionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectionsServer will
// result in compilation errors.
type UnsafeElectionsServer interface {
	mustEmbedUnimplementedElectionsServer()
}

func RegisterElectionsServer(s grpc.ServiceRegistrar, srv ElectionsServer) {
	s.RegisterService(&Elections_ServiceDesc, srv)
}

func _Elections_SubmitVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectionsServer).SubmitVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/elections_with_admin.Elections/SubmitVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectionsServer).SubmitVote(ctx, req.(*Vote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Elections_Internal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ElectionsServer).Internal(&electionsInternalServer{stream})
}

type Elections_InternalServer interface {
	Send(*StatsVote) error
	Recv() (*Vote, error)
	grpc.ServerStream
}

type electionsInternalServer struct {
	grpc.ServerStream
}

func (x *electionsInternalServer) Send(m *StatsVote) error {
	return x.ServerStream.SendMsg(m)
}

func (x *electionsInternalServer) Recv() (*Vote, error) {
	m := new(Vote)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Elections_ServiceDesc is the grpc.ServiceDesc for Elections service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Elections_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "elections_with_admin.Elections",
	HandlerType: (*ElectionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitVote",
			Handler:    _Elections_SubmitVote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Internal",
			Handler:       _Elections_Internal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "elections.proto",
}
