// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: boost_server.proto

package im_repo_boost

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	BoostService_BoostPost_FullMethodName       = "/boost.BoostService/BoostPost"
	BoostService_GetBoostedPosts_FullMethodName = "/boost.BoostService/GetBoostedPosts"
	BoostService_GetBoostStats_FullMethodName   = "/boost.BoostService/GetBoostStats"
)

// BoostServiceClient is the client API for BoostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BoostServiceClient interface {
	BoostPost(ctx context.Context, in *BoostRequest, opts ...grpc.CallOption) (*BoostResponse, error)
	GetBoostedPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BoostedPostsResponse, error)
	GetBoostStats(ctx context.Context, in *BoostStatsRequest, opts ...grpc.CallOption) (*BoostStatsResponse, error)
}

type boostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBoostServiceClient(cc grpc.ClientConnInterface) BoostServiceClient {
	return &boostServiceClient{cc}
}

func (c *boostServiceClient) BoostPost(ctx context.Context, in *BoostRequest, opts ...grpc.CallOption) (*BoostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BoostResponse)
	err := c.cc.Invoke(ctx, BoostService_BoostPost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boostServiceClient) GetBoostedPosts(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*BoostedPostsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BoostedPostsResponse)
	err := c.cc.Invoke(ctx, BoostService_GetBoostedPosts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *boostServiceClient) GetBoostStats(ctx context.Context, in *BoostStatsRequest, opts ...grpc.CallOption) (*BoostStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BoostStatsResponse)
	err := c.cc.Invoke(ctx, BoostService_GetBoostStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BoostServiceServer is the server API for BoostService service.
// All implementations must embed UnimplementedBoostServiceServer
// for forward compatibility.
type BoostServiceServer interface {
	BoostPost(context.Context, *BoostRequest) (*BoostResponse, error)
	GetBoostedPosts(context.Context, *EmptyRequest) (*BoostedPostsResponse, error)
	GetBoostStats(context.Context, *BoostStatsRequest) (*BoostStatsResponse, error)
	mustEmbedUnimplementedBoostServiceServer()
}

// UnimplementedBoostServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBoostServiceServer struct{}

func (UnimplementedBoostServiceServer) BoostPost(context.Context, *BoostRequest) (*BoostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BoostPost not implemented")
}
func (UnimplementedBoostServiceServer) GetBoostedPosts(context.Context, *EmptyRequest) (*BoostedPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoostedPosts not implemented")
}
func (UnimplementedBoostServiceServer) GetBoostStats(context.Context, *BoostStatsRequest) (*BoostStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoostStats not implemented")
}
func (UnimplementedBoostServiceServer) mustEmbedUnimplementedBoostServiceServer() {}
func (UnimplementedBoostServiceServer) testEmbeddedByValue()                      {}

// UnsafeBoostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BoostServiceServer will
// result in compilation errors.
type UnsafeBoostServiceServer interface {
	mustEmbedUnimplementedBoostServiceServer()
}

func RegisterBoostServiceServer(s grpc.ServiceRegistrar, srv BoostServiceServer) {
	// If the following call pancis, it indicates UnimplementedBoostServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BoostService_ServiceDesc, srv)
}

func _BoostService_BoostPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoostServiceServer).BoostPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoostService_BoostPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoostServiceServer).BoostPost(ctx, req.(*BoostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoostService_GetBoostedPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoostServiceServer).GetBoostedPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoostService_GetBoostedPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoostServiceServer).GetBoostedPosts(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BoostService_GetBoostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BoostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BoostServiceServer).GetBoostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BoostService_GetBoostStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BoostServiceServer).GetBoostStats(ctx, req.(*BoostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BoostService_ServiceDesc is the grpc.ServiceDesc for BoostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BoostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "boost.BoostService",
	HandlerType: (*BoostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BoostPost",
			Handler:    _BoostService_BoostPost_Handler,
		},
		{
			MethodName: "GetBoostedPosts",
			Handler:    _BoostService_GetBoostedPosts_Handler,
		},
		{
			MethodName: "GetBoostStats",
			Handler:    _BoostService_GetBoostStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "boost_server.proto",
}
