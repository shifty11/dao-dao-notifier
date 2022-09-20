// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package subscription_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	GetSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error)
	ToggleSubscription(ctx context.Context, in *ToggleSubscriptionRequest, opts ...grpc.CallOption) (*ToggleSubscriptionResponse, error)
	AddDao(ctx context.Context, in *AddDaoRequest, opts ...grpc.CallOption) (SubscriptionService_AddDaoClient, error)
	DeleteDao(ctx context.Context, in *DeleteDaoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) GetSubscriptions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSubscriptionsResponse, error) {
	out := new(GetSubscriptionsResponse)
	err := c.cc.Invoke(ctx, "/daodao_notifier_grpc.SubscriptionService/GetSubscriptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) ToggleSubscription(ctx context.Context, in *ToggleSubscriptionRequest, opts ...grpc.CallOption) (*ToggleSubscriptionResponse, error) {
	out := new(ToggleSubscriptionResponse)
	err := c.cc.Invoke(ctx, "/daodao_notifier_grpc.SubscriptionService/ToggleSubscription", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) AddDao(ctx context.Context, in *AddDaoRequest, opts ...grpc.CallOption) (SubscriptionService_AddDaoClient, error) {
	stream, err := c.cc.NewStream(ctx, &SubscriptionService_ServiceDesc.Streams[0], "/daodao_notifier_grpc.SubscriptionService/AddDao", opts...)
	if err != nil {
		return nil, err
	}
	x := &subscriptionServiceAddDaoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SubscriptionService_AddDaoClient interface {
	Recv() (*AddDaoResponse, error)
	grpc.ClientStream
}

type subscriptionServiceAddDaoClient struct {
	grpc.ClientStream
}

func (x *subscriptionServiceAddDaoClient) Recv() (*AddDaoResponse, error) {
	m := new(AddDaoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *subscriptionServiceClient) DeleteDao(ctx context.Context, in *DeleteDaoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/daodao_notifier_grpc.SubscriptionService/DeleteDao", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility
type SubscriptionServiceServer interface {
	GetSubscriptions(context.Context, *emptypb.Empty) (*GetSubscriptionsResponse, error)
	ToggleSubscription(context.Context, *ToggleSubscriptionRequest) (*ToggleSubscriptionResponse, error)
	AddDao(*AddDaoRequest, SubscriptionService_AddDaoServer) error
	DeleteDao(context.Context, *DeleteDaoRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionServiceServer struct {
}

func (UnimplementedSubscriptionServiceServer) GetSubscriptions(context.Context, *emptypb.Empty) (*GetSubscriptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptions not implemented")
}
func (UnimplementedSubscriptionServiceServer) ToggleSubscription(context.Context, *ToggleSubscriptionRequest) (*ToggleSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ToggleSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) AddDao(*AddDaoRequest, SubscriptionService_AddDaoServer) error {
	return status.Errorf(codes.Unimplemented, "method AddDao not implemented")
}
func (UnimplementedSubscriptionServiceServer) DeleteDao(context.Context, *DeleteDaoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDao not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_GetSubscriptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daodao_notifier_grpc.SubscriptionService/GetSubscriptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscriptions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_ToggleSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ToggleSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).ToggleSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daodao_notifier_grpc.SubscriptionService/ToggleSubscription",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).ToggleSubscription(ctx, req.(*ToggleSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_AddDao_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AddDaoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SubscriptionServiceServer).AddDao(m, &subscriptionServiceAddDaoServer{stream})
}

type SubscriptionService_AddDaoServer interface {
	Send(*AddDaoResponse) error
	grpc.ServerStream
}

type subscriptionServiceAddDaoServer struct {
	grpc.ServerStream
}

func (x *subscriptionServiceAddDaoServer) Send(m *AddDaoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SubscriptionService_DeleteDao_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDaoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).DeleteDao(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/daodao_notifier_grpc.SubscriptionService/DeleteDao",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).DeleteDao(ctx, req.(*DeleteDaoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "daodao_notifier_grpc.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubscriptions",
			Handler:    _SubscriptionService_GetSubscriptions_Handler,
		},
		{
			MethodName: "ToggleSubscription",
			Handler:    _SubscriptionService_ToggleSubscription_Handler,
		},
		{
			MethodName: "DeleteDao",
			Handler:    _SubscriptionService_DeleteDao_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddDao",
			Handler:       _SubscriptionService_AddDao_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "subscription_service.proto",
}
