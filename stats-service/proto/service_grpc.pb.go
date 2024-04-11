// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/service.proto

package proto

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

const (
	StatsService_FetchStats_FullMethodName = "/StatsService/FetchStats"
)

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsServiceClient interface {
	FetchStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StatsService_FetchStatsClient, error)
}

type statsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsServiceClient(cc grpc.ClientConnInterface) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) FetchStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (StatsService_FetchStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &StatsService_ServiceDesc.Streams[0], StatsService_FetchStats_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &statsServiceFetchStatsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StatsService_FetchStatsClient interface {
	Recv() (*StatsResponse, error)
	grpc.ClientStream
}

type statsServiceFetchStatsClient struct {
	grpc.ClientStream
}

func (x *statsServiceFetchStatsClient) Recv() (*StatsResponse, error) {
	m := new(StatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatsServiceServer is the server API for StatsService service.
// All implementations must embed UnimplementedStatsServiceServer
// for forward compatibility
type StatsServiceServer interface {
	FetchStats(*empty.Empty, StatsService_FetchStatsServer) error
	mustEmbedUnimplementedStatsServiceServer()
}

// UnimplementedStatsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (UnimplementedStatsServiceServer) FetchStats(*empty.Empty, StatsService_FetchStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchStats not implemented")
}
func (UnimplementedStatsServiceServer) mustEmbedUnimplementedStatsServiceServer() {}

// UnsafeStatsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServiceServer will
// result in compilation errors.
type UnsafeStatsServiceServer interface {
	mustEmbedUnimplementedStatsServiceServer()
}

func RegisterStatsServiceServer(s grpc.ServiceRegistrar, srv StatsServiceServer) {
	s.RegisterService(&StatsService_ServiceDesc, srv)
}

func _StatsService_FetchStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StatsServiceServer).FetchStats(m, &statsServiceFetchStatsServer{stream})
}

type StatsService_FetchStatsServer interface {
	Send(*StatsResponse) error
	grpc.ServerStream
}

type statsServiceFetchStatsServer struct {
	grpc.ServerStream
}

func (x *statsServiceFetchStatsServer) Send(m *StatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// StatsService_ServiceDesc is the grpc.ServiceDesc for StatsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchStats",
			Handler:       _StatsService_FetchStats_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/service.proto",
}
