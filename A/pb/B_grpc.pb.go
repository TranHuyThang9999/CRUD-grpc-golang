// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: B.proto

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	List(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*ReposeList, error)
	FindByEmail(ctx context.Context, in *RequestFindByEmail, opts ...grpc.CallOption) (*ReposeFindByEmail, error)
	Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*User, error)
	UpdateByEmail(ctx context.Context, in *RequestUpdateByEmail, opts ...grpc.CallOption) (*ResponseUpdateByEmail, error)
	DeleteByEmail(ctx context.Context, in *RequestDeleteByEmail, opts ...grpc.CallOption) (*ReposeDeleteByEmail, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) List(ctx context.Context, in *RequestList, opts ...grpc.CallOption) (*ReposeList, error) {
	out := new(ReposeList)
	err := c.cc.Invoke(ctx, "/User.UserService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FindByEmail(ctx context.Context, in *RequestFindByEmail, opts ...grpc.CallOption) (*ReposeFindByEmail, error) {
	out := new(ReposeFindByEmail)
	err := c.cc.Invoke(ctx, "/User.UserService/FindByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Create(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/User.UserService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateByEmail(ctx context.Context, in *RequestUpdateByEmail, opts ...grpc.CallOption) (*ResponseUpdateByEmail, error) {
	out := new(ResponseUpdateByEmail)
	err := c.cc.Invoke(ctx, "/User.UserService/UpdateByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteByEmail(ctx context.Context, in *RequestDeleteByEmail, opts ...grpc.CallOption) (*ReposeDeleteByEmail, error) {
	out := new(ReposeDeleteByEmail)
	err := c.cc.Invoke(ctx, "/User.UserService/DeleteByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	List(context.Context, *RequestList) (*ReposeList, error)
	FindByEmail(context.Context, *RequestFindByEmail) (*ReposeFindByEmail, error)
	Create(context.Context, *RequestCreate) (*User, error)
	UpdateByEmail(context.Context, *RequestUpdateByEmail) (*ResponseUpdateByEmail, error)
	DeleteByEmail(context.Context, *RequestDeleteByEmail) (*ReposeDeleteByEmail, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) List(context.Context, *RequestList) (*ReposeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedUserServiceServer) FindByEmail(context.Context, *RequestFindByEmail) (*ReposeFindByEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByEmail not implemented")
}
func (UnimplementedUserServiceServer) Create(context.Context, *RequestCreate) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserServiceServer) UpdateByEmail(context.Context, *RequestUpdateByEmail) (*ResponseUpdateByEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByEmail not implemented")
}
func (UnimplementedUserServiceServer) DeleteByEmail(context.Context, *RequestDeleteByEmail) (*ReposeDeleteByEmail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByEmail not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).List(ctx, req.(*RequestList))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FindByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFindByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FindByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/FindByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FindByEmail(ctx, req.(*RequestFindByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Create(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUpdateByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/UpdateByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateByEmail(ctx, req.(*RequestUpdateByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeleteByEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/User.UserService/DeleteByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteByEmail(ctx, req.(*RequestDeleteByEmail))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "User.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _UserService_List_Handler,
		},
		{
			MethodName: "FindByEmail",
			Handler:    _UserService_FindByEmail_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _UserService_Create_Handler,
		},
		{
			MethodName: "UpdateByEmail",
			Handler:    _UserService_UpdateByEmail_Handler,
		},
		{
			MethodName: "DeleteByEmail",
			Handler:    _UserService_DeleteByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "B.proto",
}
