// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.6
// source: api.proto

package pb

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

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
	SaveStats(ctx context.Context, in *SaveStatsRequest, opts ...grpc.CallOption) (*SaveStatsResponse, error)
	SaveStatsStream(ctx context.Context, opts ...grpc.CallOption) (Stats_SaveStatsStreamClient, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/proto.Stats/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) SaveStats(ctx context.Context, in *SaveStatsRequest, opts ...grpc.CallOption) (*SaveStatsResponse, error) {
	out := new(SaveStatsResponse)
	err := c.cc.Invoke(ctx, "/proto.Stats/SaveStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) SaveStatsStream(ctx context.Context, opts ...grpc.CallOption) (Stats_SaveStatsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stats_ServiceDesc.Streams[0], "/proto.Stats/SaveStatsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &statsSaveStatsStreamClient{stream}
	return x, nil
}

type Stats_SaveStatsStreamClient interface {
	Send(*FooData) error
	CloseAndRecv() (*SaveStatsResponse, error)
	grpc.ClientStream
}

type statsSaveStatsStreamClient struct {
	grpc.ClientStream
}

func (x *statsSaveStatsStreamClient) Send(m *FooData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *statsSaveStatsStreamClient) CloseAndRecv() (*SaveStatsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SaveStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility
type StatsServer interface {
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
	SaveStats(context.Context, *SaveStatsRequest) (*SaveStatsResponse, error)
	SaveStatsStream(Stats_SaveStatsStreamServer) error
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (UnimplementedStatsServer) GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedStatsServer) SaveStats(context.Context, *SaveStatsRequest) (*SaveStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveStats not implemented")
}
func (UnimplementedStatsServer) SaveStatsStream(Stats_SaveStatsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SaveStatsStream not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Stats/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_SaveStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).SaveStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Stats/SaveStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).SaveStats(ctx, req.(*SaveStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_SaveStatsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StatsServer).SaveStatsStream(&statsSaveStatsStreamServer{stream})
}

type Stats_SaveStatsStreamServer interface {
	SendAndClose(*SaveStatsResponse) error
	Recv() (*FooData, error)
	grpc.ServerStream
}

type statsSaveStatsStreamServer struct {
	grpc.ServerStream
}

func (x *statsSaveStatsStreamServer) SendAndClose(m *SaveStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *statsSaveStatsStreamServer) Recv() (*FooData, error) {
	m := new(FooData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _Stats_GetStats_Handler,
		},
		{
			MethodName: "SaveStats",
			Handler:    _Stats_SaveStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SaveStatsStream",
			Handler:       _Stats_SaveStatsStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api.proto",
}
