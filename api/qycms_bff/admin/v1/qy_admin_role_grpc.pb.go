// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_role.proto

package v1

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

// QyAdminRoleClient is the client API for QyAdminRole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminRoleClient interface {
	CreateQyAdminRole(ctx context.Context, in *CreateQyAdminRoleRequest, opts ...grpc.CallOption) (*CreateQyAdminRoleReply, error)
	UpdateQyAdminRole(ctx context.Context, in *UpdateQyAdminRoleRequest, opts ...grpc.CallOption) (*UpdateQyAdminRoleReply, error)
	DeleteQyAdminRole(ctx context.Context, in *DeleteQyAdminRoleRequest, opts ...grpc.CallOption) (*DeleteQyAdminRoleReply, error)
	DeleteQyAdminRoles(ctx context.Context, in *DeleteQyAdminRolesRequest, opts ...grpc.CallOption) (*DeleteQyAdminRolesReply, error)
	GetQyAdminRole(ctx context.Context, in *GetQyAdminRoleRequest, opts ...grpc.CallOption) (*GetQyAdminRoleReply, error)
	ListQyAdminRole(ctx context.Context, in *ListQyAdminRoleRequest, opts ...grpc.CallOption) (*ListQyAdminRoleReply, error)
	SaveQyAdminRoleMenus(ctx context.Context, in *SaveRoleMenusRequest, opts ...grpc.CallOption) (*SaveRoleMenusReply, error)
	SaveQyAdminRoleApis(ctx context.Context, in *SaveRoleApisRequest, opts ...grpc.CallOption) (*SaveRoleApisReply, error)
}

type qyAdminRoleClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminRoleClient(cc grpc.ClientConnInterface) QyAdminRoleClient {
	return &qyAdminRoleClient{cc}
}

func (c *qyAdminRoleClient) CreateQyAdminRole(ctx context.Context, in *CreateQyAdminRoleRequest, opts ...grpc.CallOption) (*CreateQyAdminRoleReply, error) {
	out := new(CreateQyAdminRoleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/CreateQyAdminRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) UpdateQyAdminRole(ctx context.Context, in *UpdateQyAdminRoleRequest, opts ...grpc.CallOption) (*UpdateQyAdminRoleReply, error) {
	out := new(UpdateQyAdminRoleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/UpdateQyAdminRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) DeleteQyAdminRole(ctx context.Context, in *DeleteQyAdminRoleRequest, opts ...grpc.CallOption) (*DeleteQyAdminRoleReply, error) {
	out := new(DeleteQyAdminRoleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/DeleteQyAdminRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) DeleteQyAdminRoles(ctx context.Context, in *DeleteQyAdminRolesRequest, opts ...grpc.CallOption) (*DeleteQyAdminRolesReply, error) {
	out := new(DeleteQyAdminRolesReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/DeleteQyAdminRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) GetQyAdminRole(ctx context.Context, in *GetQyAdminRoleRequest, opts ...grpc.CallOption) (*GetQyAdminRoleReply, error) {
	out := new(GetQyAdminRoleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/GetQyAdminRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) ListQyAdminRole(ctx context.Context, in *ListQyAdminRoleRequest, opts ...grpc.CallOption) (*ListQyAdminRoleReply, error) {
	out := new(ListQyAdminRoleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/ListQyAdminRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) SaveQyAdminRoleMenus(ctx context.Context, in *SaveRoleMenusRequest, opts ...grpc.CallOption) (*SaveRoleMenusReply, error) {
	out := new(SaveRoleMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/SaveQyAdminRoleMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminRoleClient) SaveQyAdminRoleApis(ctx context.Context, in *SaveRoleApisRequest, opts ...grpc.CallOption) (*SaveRoleApisReply, error) {
	out := new(SaveRoleApisReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminRole/SaveQyAdminRoleApis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminRoleServer is the server API for QyAdminRole service.
// All implementations must embed UnimplementedQyAdminRoleServer
// for forward compatibility
type QyAdminRoleServer interface {
	CreateQyAdminRole(context.Context, *CreateQyAdminRoleRequest) (*CreateQyAdminRoleReply, error)
	UpdateQyAdminRole(context.Context, *UpdateQyAdminRoleRequest) (*UpdateQyAdminRoleReply, error)
	DeleteQyAdminRole(context.Context, *DeleteQyAdminRoleRequest) (*DeleteQyAdminRoleReply, error)
	DeleteQyAdminRoles(context.Context, *DeleteQyAdminRolesRequest) (*DeleteQyAdminRolesReply, error)
	GetQyAdminRole(context.Context, *GetQyAdminRoleRequest) (*GetQyAdminRoleReply, error)
	ListQyAdminRole(context.Context, *ListQyAdminRoleRequest) (*ListQyAdminRoleReply, error)
	SaveQyAdminRoleMenus(context.Context, *SaveRoleMenusRequest) (*SaveRoleMenusReply, error)
	SaveQyAdminRoleApis(context.Context, *SaveRoleApisRequest) (*SaveRoleApisReply, error)
	mustEmbedUnimplementedQyAdminRoleServer()
}

// UnimplementedQyAdminRoleServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminRoleServer struct {
}

func (UnimplementedQyAdminRoleServer) CreateQyAdminRole(context.Context, *CreateQyAdminRoleRequest) (*CreateQyAdminRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminRole not implemented")
}
func (UnimplementedQyAdminRoleServer) UpdateQyAdminRole(context.Context, *UpdateQyAdminRoleRequest) (*UpdateQyAdminRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminRole not implemented")
}
func (UnimplementedQyAdminRoleServer) DeleteQyAdminRole(context.Context, *DeleteQyAdminRoleRequest) (*DeleteQyAdminRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminRole not implemented")
}
func (UnimplementedQyAdminRoleServer) DeleteQyAdminRoles(context.Context, *DeleteQyAdminRolesRequest) (*DeleteQyAdminRolesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminRoles not implemented")
}
func (UnimplementedQyAdminRoleServer) GetQyAdminRole(context.Context, *GetQyAdminRoleRequest) (*GetQyAdminRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminRole not implemented")
}
func (UnimplementedQyAdminRoleServer) ListQyAdminRole(context.Context, *ListQyAdminRoleRequest) (*ListQyAdminRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminRole not implemented")
}
func (UnimplementedQyAdminRoleServer) SaveQyAdminRoleMenus(context.Context, *SaveRoleMenusRequest) (*SaveRoleMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQyAdminRoleMenus not implemented")
}
func (UnimplementedQyAdminRoleServer) SaveQyAdminRoleApis(context.Context, *SaveRoleApisRequest) (*SaveRoleApisReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveQyAdminRoleApis not implemented")
}
func (UnimplementedQyAdminRoleServer) mustEmbedUnimplementedQyAdminRoleServer() {}

// UnsafeQyAdminRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminRoleServer will
// result in compilation errors.
type UnsafeQyAdminRoleServer interface {
	mustEmbedUnimplementedQyAdminRoleServer()
}

func RegisterQyAdminRoleServer(s grpc.ServiceRegistrar, srv QyAdminRoleServer) {
	s.RegisterService(&QyAdminRole_ServiceDesc, srv)
}

func _QyAdminRole_CreateQyAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).CreateQyAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/CreateQyAdminRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).CreateQyAdminRole(ctx, req.(*CreateQyAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_UpdateQyAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).UpdateQyAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/UpdateQyAdminRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).UpdateQyAdminRole(ctx, req.(*UpdateQyAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_DeleteQyAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).DeleteQyAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/DeleteQyAdminRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).DeleteQyAdminRole(ctx, req.(*DeleteQyAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_DeleteQyAdminRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).DeleteQyAdminRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/DeleteQyAdminRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).DeleteQyAdminRoles(ctx, req.(*DeleteQyAdminRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_GetQyAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).GetQyAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/GetQyAdminRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).GetQyAdminRole(ctx, req.(*GetQyAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_ListQyAdminRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).ListQyAdminRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/ListQyAdminRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).ListQyAdminRole(ctx, req.(*ListQyAdminRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_SaveQyAdminRoleMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRoleMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).SaveQyAdminRoleMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/SaveQyAdminRoleMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).SaveQyAdminRoleMenus(ctx, req.(*SaveRoleMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminRole_SaveQyAdminRoleApis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRoleApisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminRoleServer).SaveQyAdminRoleApis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminRole/SaveQyAdminRoleApis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminRoleServer).SaveQyAdminRoleApis(ctx, req.(*SaveRoleApisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminRole_ServiceDesc is the grpc.ServiceDesc for QyAdminRole service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminRole_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminRole",
	HandlerType: (*QyAdminRoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminRole",
			Handler:    _QyAdminRole_CreateQyAdminRole_Handler,
		},
		{
			MethodName: "UpdateQyAdminRole",
			Handler:    _QyAdminRole_UpdateQyAdminRole_Handler,
		},
		{
			MethodName: "DeleteQyAdminRole",
			Handler:    _QyAdminRole_DeleteQyAdminRole_Handler,
		},
		{
			MethodName: "DeleteQyAdminRoles",
			Handler:    _QyAdminRole_DeleteQyAdminRoles_Handler,
		},
		{
			MethodName: "GetQyAdminRole",
			Handler:    _QyAdminRole_GetQyAdminRole_Handler,
		},
		{
			MethodName: "ListQyAdminRole",
			Handler:    _QyAdminRole_ListQyAdminRole_Handler,
		},
		{
			MethodName: "SaveQyAdminRoleMenus",
			Handler:    _QyAdminRole_SaveQyAdminRoleMenus_Handler,
		},
		{
			MethodName: "SaveQyAdminRoleApis",
			Handler:    _QyAdminRole_SaveQyAdminRoleApis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_role.proto",
}
