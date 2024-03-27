// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user_basic_service.proto

package user_basic

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
	UserBasicService_UpdateUserInfo_FullMethodName = "/user_basic.service.v1.UserBasicService/UpdateUserInfo"
)

// UserBasicServiceClient is the client API for UserBasicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserBasicServiceClient interface {
	UpdateUserInfo(ctx context.Context, in *MemberMessage, opts ...grpc.CallOption) (*MemberMessage, error)
}

type userBasicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserBasicServiceClient(cc grpc.ClientConnInterface) UserBasicServiceClient {
	return &userBasicServiceClient{cc}
}

func (c *userBasicServiceClient) UpdateUserInfo(ctx context.Context, in *MemberMessage, opts ...grpc.CallOption) (*MemberMessage, error) {
	out := new(MemberMessage)
	err := c.cc.Invoke(ctx, UserBasicService_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserBasicServiceServer is the server API for UserBasicService service.
// All implementations must embed UnimplementedUserBasicServiceServer
// for forward compatibility
type UserBasicServiceServer interface {
	UpdateUserInfo(context.Context, *MemberMessage) (*MemberMessage, error)
	mustEmbedUnimplementedUserBasicServiceServer()
}

// UnimplementedUserBasicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserBasicServiceServer struct {
}

func (UnimplementedUserBasicServiceServer) UpdateUserInfo(context.Context, *MemberMessage) (*MemberMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserBasicServiceServer) mustEmbedUnimplementedUserBasicServiceServer() {}

// UnsafeUserBasicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserBasicServiceServer will
// result in compilation errors.
type UnsafeUserBasicServiceServer interface {
	mustEmbedUnimplementedUserBasicServiceServer()
}

func RegisterUserBasicServiceServer(s grpc.ServiceRegistrar, srv UserBasicServiceServer) {
	s.RegisterService(&UserBasicService_ServiceDesc, srv)
}

func _UserBasicService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserBasicServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserBasicService_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserBasicServiceServer).UpdateUserInfo(ctx, req.(*MemberMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// UserBasicService_ServiceDesc is the grpc.ServiceDesc for UserBasicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserBasicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_basic.service.v1.UserBasicService",
	HandlerType: (*UserBasicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserBasicService_UpdateUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_basic_service.proto",
}
