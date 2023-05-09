// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.3
// source: system.proto

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

const (
	SystemRpc_RetrieveNotice_FullMethodName    = "/system.system_rpc/RetrieveNotice"
	SystemRpc_ListNotice_FullMethodName        = "/system.system_rpc/ListNotice"
	SystemRpc_CreateNotice_FullMethodName      = "/system.system_rpc/CreateNotice"
	SystemRpc_UpdateNotice_FullMethodName      = "/system.system_rpc/UpdateNotice"
	SystemRpc_DeleteNotice_FullMethodName      = "/system.system_rpc/DeleteNotice"
	SystemRpc_RetrieveNews_FullMethodName      = "/system.system_rpc/RetrieveNews"
	SystemRpc_ListNews_FullMethodName          = "/system.system_rpc/ListNews"
	SystemRpc_CreateNews_FullMethodName        = "/system.system_rpc/CreateNews"
	SystemRpc_UpdateNews_FullMethodName        = "/system.system_rpc/UpdateNews"
	SystemRpc_DeleteNews_FullMethodName        = "/system.system_rpc/DeleteNews"
	SystemRpc_RetrieveAdmin_FullMethodName     = "/system.system_rpc/RetrieveAdmin"
	SystemRpc_ListAdmin_FullMethodName         = "/system.system_rpc/ListAdmin"
	SystemRpc_CreateAdmin_FullMethodName       = "/system.system_rpc/CreateAdmin"
	SystemRpc_UpdateAdmin_FullMethodName       = "/system.system_rpc/UpdateAdmin"
	SystemRpc_DeleteAdmin_FullMethodName       = "/system.system_rpc/DeleteAdmin"
	SystemRpc_RetrieveUserRole_FullMethodName  = "/system.system_rpc/RetrieveUserRole"
	SystemRpc_UpdateUserRole_FullMethodName    = "/system.system_rpc/UpdateUserRole"
	SystemRpc_ContainsRole_FullMethodName      = "/system.system_rpc/ContainsRole"
	SystemRpc_RetrieveCommunity_FullMethodName = "/system.system_rpc/RetrieveCommunity"
	SystemRpc_ListCommunity_FullMethodName     = "/system.system_rpc/ListCommunity"
	SystemRpc_CreateCommunity_FullMethodName   = "/system.system_rpc/CreateCommunity"
	SystemRpc_UpdateCommunity_FullMethodName   = "/system.system_rpc/UpdateCommunity"
	SystemRpc_DeleteCommunity_FullMethodName   = "/system.system_rpc/DeleteCommunity"
	SystemRpc_ListUseridByRole_FullMethodName  = "/system.system_rpc/ListUseridByRole"
)

// SystemRpcClient is the client API for SystemRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SystemRpcClient interface {
	RetrieveNotice(ctx context.Context, in *RetrieveNoticeReq, opts ...grpc.CallOption) (*RetrieveNoticeResp, error)
	ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error)
	CreateNotice(ctx context.Context, in *CreateNoticeReq, opts ...grpc.CallOption) (*CreateNoticeResp, error)
	UpdateNotice(ctx context.Context, in *UpdateNoticeReq, opts ...grpc.CallOption) (*UpdateNoticeResp, error)
	DeleteNotice(ctx context.Context, in *DeleteNoticeReq, opts ...grpc.CallOption) (*DeleteNoticeResp, error)
	RetrieveNews(ctx context.Context, in *RetrieveNewsReq, opts ...grpc.CallOption) (*RetrieveNewsResp, error)
	ListNews(ctx context.Context, in *ListNewsReq, opts ...grpc.CallOption) (*ListNewsResp, error)
	CreateNews(ctx context.Context, in *CreateNewsReq, opts ...grpc.CallOption) (*CreateNewsResp, error)
	UpdateNews(ctx context.Context, in *UpdateNewsReq, opts ...grpc.CallOption) (*UpdateNewsResp, error)
	DeleteNews(ctx context.Context, in *DeleteNewsReq, opts ...grpc.CallOption) (*DeleteNewsResp, error)
	RetrieveAdmin(ctx context.Context, in *RetrieveAdminReq, opts ...grpc.CallOption) (*RetrieveAdminResp, error)
	ListAdmin(ctx context.Context, in *ListAdminReq, opts ...grpc.CallOption) (*ListAdminResp, error)
	CreateAdmin(ctx context.Context, in *CreateAdminReq, opts ...grpc.CallOption) (*CreateAdminResp, error)
	UpdateAdmin(ctx context.Context, in *UpdateAdminReq, opts ...grpc.CallOption) (*UpdateAdminResp, error)
	DeleteAdmin(ctx context.Context, in *DeleteAdminReq, opts ...grpc.CallOption) (*DeleteAdminResp, error)
	// 获取用户的所有角色
	RetrieveUserRole(ctx context.Context, in *RetrieveUserRoleReq, opts ...grpc.CallOption) (*RetrieveUserRoleResp, error)
	// 更新用户的角色
	UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleResp, error)
	ContainsRole(ctx context.Context, in *ContainsRoleReq, opts ...grpc.CallOption) (*ContainsRoleResp, error)
	RetrieveCommunity(ctx context.Context, in *RetrieveCommunityReq, opts ...grpc.CallOption) (*RetrieveCommunityResp, error)
	ListCommunity(ctx context.Context, in *ListCommunityReq, opts ...grpc.CallOption) (*ListCommunityResp, error)
	CreateCommunity(ctx context.Context, in *CreateCommunityReq, opts ...grpc.CallOption) (*CreateCommunityResp, error)
	UpdateCommunity(ctx context.Context, in *UpdateCommunityReq, opts ...grpc.CallOption) (*UpdateCommunityResp, error)
	DeleteCommunity(ctx context.Context, in *DeleteCommunityReq, opts ...grpc.CallOption) (*DeleteCommunityResp, error)
	ListUseridByRole(ctx context.Context, in *ListUseridReq, opts ...grpc.CallOption) (*ListUseridResp, error)
}

type systemRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSystemRpcClient(cc grpc.ClientConnInterface) SystemRpcClient {
	return &systemRpcClient{cc}
}

func (c *systemRpcClient) RetrieveNotice(ctx context.Context, in *RetrieveNoticeReq, opts ...grpc.CallOption) (*RetrieveNoticeResp, error) {
	out := new(RetrieveNoticeResp)
	err := c.cc.Invoke(ctx, SystemRpc_RetrieveNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error) {
	out := new(ListNoticeResp)
	err := c.cc.Invoke(ctx, SystemRpc_ListNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) CreateNotice(ctx context.Context, in *CreateNoticeReq, opts ...grpc.CallOption) (*CreateNoticeResp, error) {
	out := new(CreateNoticeResp)
	err := c.cc.Invoke(ctx, SystemRpc_CreateNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) UpdateNotice(ctx context.Context, in *UpdateNoticeReq, opts ...grpc.CallOption) (*UpdateNoticeResp, error) {
	out := new(UpdateNoticeResp)
	err := c.cc.Invoke(ctx, SystemRpc_UpdateNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) DeleteNotice(ctx context.Context, in *DeleteNoticeReq, opts ...grpc.CallOption) (*DeleteNoticeResp, error) {
	out := new(DeleteNoticeResp)
	err := c.cc.Invoke(ctx, SystemRpc_DeleteNotice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) RetrieveNews(ctx context.Context, in *RetrieveNewsReq, opts ...grpc.CallOption) (*RetrieveNewsResp, error) {
	out := new(RetrieveNewsResp)
	err := c.cc.Invoke(ctx, SystemRpc_RetrieveNews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ListNews(ctx context.Context, in *ListNewsReq, opts ...grpc.CallOption) (*ListNewsResp, error) {
	out := new(ListNewsResp)
	err := c.cc.Invoke(ctx, SystemRpc_ListNews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) CreateNews(ctx context.Context, in *CreateNewsReq, opts ...grpc.CallOption) (*CreateNewsResp, error) {
	out := new(CreateNewsResp)
	err := c.cc.Invoke(ctx, SystemRpc_CreateNews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) UpdateNews(ctx context.Context, in *UpdateNewsReq, opts ...grpc.CallOption) (*UpdateNewsResp, error) {
	out := new(UpdateNewsResp)
	err := c.cc.Invoke(ctx, SystemRpc_UpdateNews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) DeleteNews(ctx context.Context, in *DeleteNewsReq, opts ...grpc.CallOption) (*DeleteNewsResp, error) {
	out := new(DeleteNewsResp)
	err := c.cc.Invoke(ctx, SystemRpc_DeleteNews_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) RetrieveAdmin(ctx context.Context, in *RetrieveAdminReq, opts ...grpc.CallOption) (*RetrieveAdminResp, error) {
	out := new(RetrieveAdminResp)
	err := c.cc.Invoke(ctx, SystemRpc_RetrieveAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ListAdmin(ctx context.Context, in *ListAdminReq, opts ...grpc.CallOption) (*ListAdminResp, error) {
	out := new(ListAdminResp)
	err := c.cc.Invoke(ctx, SystemRpc_ListAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) CreateAdmin(ctx context.Context, in *CreateAdminReq, opts ...grpc.CallOption) (*CreateAdminResp, error) {
	out := new(CreateAdminResp)
	err := c.cc.Invoke(ctx, SystemRpc_CreateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) UpdateAdmin(ctx context.Context, in *UpdateAdminReq, opts ...grpc.CallOption) (*UpdateAdminResp, error) {
	out := new(UpdateAdminResp)
	err := c.cc.Invoke(ctx, SystemRpc_UpdateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) DeleteAdmin(ctx context.Context, in *DeleteAdminReq, opts ...grpc.CallOption) (*DeleteAdminResp, error) {
	out := new(DeleteAdminResp)
	err := c.cc.Invoke(ctx, SystemRpc_DeleteAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) RetrieveUserRole(ctx context.Context, in *RetrieveUserRoleReq, opts ...grpc.CallOption) (*RetrieveUserRoleResp, error) {
	out := new(RetrieveUserRoleResp)
	err := c.cc.Invoke(ctx, SystemRpc_RetrieveUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleResp, error) {
	out := new(UpdateUserRoleResp)
	err := c.cc.Invoke(ctx, SystemRpc_UpdateUserRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ContainsRole(ctx context.Context, in *ContainsRoleReq, opts ...grpc.CallOption) (*ContainsRoleResp, error) {
	out := new(ContainsRoleResp)
	err := c.cc.Invoke(ctx, SystemRpc_ContainsRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) RetrieveCommunity(ctx context.Context, in *RetrieveCommunityReq, opts ...grpc.CallOption) (*RetrieveCommunityResp, error) {
	out := new(RetrieveCommunityResp)
	err := c.cc.Invoke(ctx, SystemRpc_RetrieveCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ListCommunity(ctx context.Context, in *ListCommunityReq, opts ...grpc.CallOption) (*ListCommunityResp, error) {
	out := new(ListCommunityResp)
	err := c.cc.Invoke(ctx, SystemRpc_ListCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) CreateCommunity(ctx context.Context, in *CreateCommunityReq, opts ...grpc.CallOption) (*CreateCommunityResp, error) {
	out := new(CreateCommunityResp)
	err := c.cc.Invoke(ctx, SystemRpc_CreateCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) UpdateCommunity(ctx context.Context, in *UpdateCommunityReq, opts ...grpc.CallOption) (*UpdateCommunityResp, error) {
	out := new(UpdateCommunityResp)
	err := c.cc.Invoke(ctx, SystemRpc_UpdateCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) DeleteCommunity(ctx context.Context, in *DeleteCommunityReq, opts ...grpc.CallOption) (*DeleteCommunityResp, error) {
	out := new(DeleteCommunityResp)
	err := c.cc.Invoke(ctx, SystemRpc_DeleteCommunity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemRpcClient) ListUseridByRole(ctx context.Context, in *ListUseridReq, opts ...grpc.CallOption) (*ListUseridResp, error) {
	out := new(ListUseridResp)
	err := c.cc.Invoke(ctx, SystemRpc_ListUseridByRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemRpcServer is the server API for SystemRpc service.
// All implementations must embed UnimplementedSystemRpcServer
// for forward compatibility
type SystemRpcServer interface {
	RetrieveNotice(context.Context, *RetrieveNoticeReq) (*RetrieveNoticeResp, error)
	ListNotice(context.Context, *ListNoticeReq) (*ListNoticeResp, error)
	CreateNotice(context.Context, *CreateNoticeReq) (*CreateNoticeResp, error)
	UpdateNotice(context.Context, *UpdateNoticeReq) (*UpdateNoticeResp, error)
	DeleteNotice(context.Context, *DeleteNoticeReq) (*DeleteNoticeResp, error)
	RetrieveNews(context.Context, *RetrieveNewsReq) (*RetrieveNewsResp, error)
	ListNews(context.Context, *ListNewsReq) (*ListNewsResp, error)
	CreateNews(context.Context, *CreateNewsReq) (*CreateNewsResp, error)
	UpdateNews(context.Context, *UpdateNewsReq) (*UpdateNewsResp, error)
	DeleteNews(context.Context, *DeleteNewsReq) (*DeleteNewsResp, error)
	RetrieveAdmin(context.Context, *RetrieveAdminReq) (*RetrieveAdminResp, error)
	ListAdmin(context.Context, *ListAdminReq) (*ListAdminResp, error)
	CreateAdmin(context.Context, *CreateAdminReq) (*CreateAdminResp, error)
	UpdateAdmin(context.Context, *UpdateAdminReq) (*UpdateAdminResp, error)
	DeleteAdmin(context.Context, *DeleteAdminReq) (*DeleteAdminResp, error)
	// 获取用户的所有角色
	RetrieveUserRole(context.Context, *RetrieveUserRoleReq) (*RetrieveUserRoleResp, error)
	// 更新用户的角色
	UpdateUserRole(context.Context, *UpdateUserRoleReq) (*UpdateUserRoleResp, error)
	ContainsRole(context.Context, *ContainsRoleReq) (*ContainsRoleResp, error)
	RetrieveCommunity(context.Context, *RetrieveCommunityReq) (*RetrieveCommunityResp, error)
	ListCommunity(context.Context, *ListCommunityReq) (*ListCommunityResp, error)
	CreateCommunity(context.Context, *CreateCommunityReq) (*CreateCommunityResp, error)
	UpdateCommunity(context.Context, *UpdateCommunityReq) (*UpdateCommunityResp, error)
	DeleteCommunity(context.Context, *DeleteCommunityReq) (*DeleteCommunityResp, error)
	ListUseridByRole(context.Context, *ListUseridReq) (*ListUseridResp, error)
	mustEmbedUnimplementedSystemRpcServer()
}

// UnimplementedSystemRpcServer must be embedded to have forward compatible implementations.
type UnimplementedSystemRpcServer struct {
}

func (UnimplementedSystemRpcServer) RetrieveNotice(context.Context, *RetrieveNoticeReq) (*RetrieveNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveNotice not implemented")
}
func (UnimplementedSystemRpcServer) ListNotice(context.Context, *ListNoticeReq) (*ListNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotice not implemented")
}
func (UnimplementedSystemRpcServer) CreateNotice(context.Context, *CreateNoticeReq) (*CreateNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNotice not implemented")
}
func (UnimplementedSystemRpcServer) UpdateNotice(context.Context, *UpdateNoticeReq) (*UpdateNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNotice not implemented")
}
func (UnimplementedSystemRpcServer) DeleteNotice(context.Context, *DeleteNoticeReq) (*DeleteNoticeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNotice not implemented")
}
func (UnimplementedSystemRpcServer) RetrieveNews(context.Context, *RetrieveNewsReq) (*RetrieveNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveNews not implemented")
}
func (UnimplementedSystemRpcServer) ListNews(context.Context, *ListNewsReq) (*ListNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNews not implemented")
}
func (UnimplementedSystemRpcServer) CreateNews(context.Context, *CreateNewsReq) (*CreateNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNews not implemented")
}
func (UnimplementedSystemRpcServer) UpdateNews(context.Context, *UpdateNewsReq) (*UpdateNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNews not implemented")
}
func (UnimplementedSystemRpcServer) DeleteNews(context.Context, *DeleteNewsReq) (*DeleteNewsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNews not implemented")
}
func (UnimplementedSystemRpcServer) RetrieveAdmin(context.Context, *RetrieveAdminReq) (*RetrieveAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveAdmin not implemented")
}
func (UnimplementedSystemRpcServer) ListAdmin(context.Context, *ListAdminReq) (*ListAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAdmin not implemented")
}
func (UnimplementedSystemRpcServer) CreateAdmin(context.Context, *CreateAdminReq) (*CreateAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdmin not implemented")
}
func (UnimplementedSystemRpcServer) UpdateAdmin(context.Context, *UpdateAdminReq) (*UpdateAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedSystemRpcServer) DeleteAdmin(context.Context, *DeleteAdminReq) (*DeleteAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAdmin not implemented")
}
func (UnimplementedSystemRpcServer) RetrieveUserRole(context.Context, *RetrieveUserRoleReq) (*RetrieveUserRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveUserRole not implemented")
}
func (UnimplementedSystemRpcServer) UpdateUserRole(context.Context, *UpdateUserRoleReq) (*UpdateUserRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedSystemRpcServer) ContainsRole(context.Context, *ContainsRoleReq) (*ContainsRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContainsRole not implemented")
}
func (UnimplementedSystemRpcServer) RetrieveCommunity(context.Context, *RetrieveCommunityReq) (*RetrieveCommunityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveCommunity not implemented")
}
func (UnimplementedSystemRpcServer) ListCommunity(context.Context, *ListCommunityReq) (*ListCommunityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommunity not implemented")
}
func (UnimplementedSystemRpcServer) CreateCommunity(context.Context, *CreateCommunityReq) (*CreateCommunityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunity not implemented")
}
func (UnimplementedSystemRpcServer) UpdateCommunity(context.Context, *UpdateCommunityReq) (*UpdateCommunityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCommunity not implemented")
}
func (UnimplementedSystemRpcServer) DeleteCommunity(context.Context, *DeleteCommunityReq) (*DeleteCommunityResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommunity not implemented")
}
func (UnimplementedSystemRpcServer) ListUseridByRole(context.Context, *ListUseridReq) (*ListUseridResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUseridByRole not implemented")
}
func (UnimplementedSystemRpcServer) mustEmbedUnimplementedSystemRpcServer() {}

// UnsafeSystemRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SystemRpcServer will
// result in compilation errors.
type UnsafeSystemRpcServer interface {
	mustEmbedUnimplementedSystemRpcServer()
}

func RegisterSystemRpcServer(s grpc.ServiceRegistrar, srv SystemRpcServer) {
	s.RegisterService(&SystemRpc_ServiceDesc, srv)
}

func _SystemRpc_RetrieveNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).RetrieveNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_RetrieveNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).RetrieveNotice(ctx, req.(*RetrieveNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ListNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ListNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ListNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ListNotice(ctx, req.(*ListNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_CreateNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).CreateNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_CreateNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).CreateNotice(ctx, req.(*CreateNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_UpdateNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).UpdateNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_UpdateNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).UpdateNotice(ctx, req.(*UpdateNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_DeleteNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).DeleteNotice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_DeleteNotice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).DeleteNotice(ctx, req.(*DeleteNoticeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_RetrieveNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).RetrieveNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_RetrieveNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).RetrieveNews(ctx, req.(*RetrieveNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ListNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ListNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ListNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ListNews(ctx, req.(*ListNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_CreateNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).CreateNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_CreateNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).CreateNews(ctx, req.(*CreateNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_UpdateNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).UpdateNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_UpdateNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).UpdateNews(ctx, req.(*UpdateNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_DeleteNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNewsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).DeleteNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_DeleteNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).DeleteNews(ctx, req.(*DeleteNewsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_RetrieveAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).RetrieveAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_RetrieveAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).RetrieveAdmin(ctx, req.(*RetrieveAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ListAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ListAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ListAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ListAdmin(ctx, req.(*ListAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_CreateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).CreateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_CreateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).CreateAdmin(ctx, req.(*CreateAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_UpdateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).UpdateAdmin(ctx, req.(*UpdateAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_DeleteAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).DeleteAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_DeleteAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).DeleteAdmin(ctx, req.(*DeleteAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_RetrieveUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).RetrieveUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_RetrieveUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).RetrieveUserRole(ctx, req.(*RetrieveUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_UpdateUserRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).UpdateUserRole(ctx, req.(*UpdateUserRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ContainsRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContainsRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ContainsRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ContainsRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ContainsRole(ctx, req.(*ContainsRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_RetrieveCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveCommunityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).RetrieveCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_RetrieveCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).RetrieveCommunity(ctx, req.(*RetrieveCommunityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ListCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommunityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ListCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ListCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ListCommunity(ctx, req.(*ListCommunityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_CreateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).CreateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_CreateCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).CreateCommunity(ctx, req.(*CreateCommunityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_UpdateCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommunityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).UpdateCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_UpdateCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).UpdateCommunity(ctx, req.(*UpdateCommunityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_DeleteCommunity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommunityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).DeleteCommunity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_DeleteCommunity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).DeleteCommunity(ctx, req.(*DeleteCommunityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemRpc_ListUseridByRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUseridReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemRpcServer).ListUseridByRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SystemRpc_ListUseridByRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemRpcServer).ListUseridByRole(ctx, req.(*ListUseridReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SystemRpc_ServiceDesc is the grpc.ServiceDesc for SystemRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SystemRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "system.system_rpc",
	HandlerType: (*SystemRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveNotice",
			Handler:    _SystemRpc_RetrieveNotice_Handler,
		},
		{
			MethodName: "ListNotice",
			Handler:    _SystemRpc_ListNotice_Handler,
		},
		{
			MethodName: "CreateNotice",
			Handler:    _SystemRpc_CreateNotice_Handler,
		},
		{
			MethodName: "UpdateNotice",
			Handler:    _SystemRpc_UpdateNotice_Handler,
		},
		{
			MethodName: "DeleteNotice",
			Handler:    _SystemRpc_DeleteNotice_Handler,
		},
		{
			MethodName: "RetrieveNews",
			Handler:    _SystemRpc_RetrieveNews_Handler,
		},
		{
			MethodName: "ListNews",
			Handler:    _SystemRpc_ListNews_Handler,
		},
		{
			MethodName: "CreateNews",
			Handler:    _SystemRpc_CreateNews_Handler,
		},
		{
			MethodName: "UpdateNews",
			Handler:    _SystemRpc_UpdateNews_Handler,
		},
		{
			MethodName: "DeleteNews",
			Handler:    _SystemRpc_DeleteNews_Handler,
		},
		{
			MethodName: "RetrieveAdmin",
			Handler:    _SystemRpc_RetrieveAdmin_Handler,
		},
		{
			MethodName: "ListAdmin",
			Handler:    _SystemRpc_ListAdmin_Handler,
		},
		{
			MethodName: "CreateAdmin",
			Handler:    _SystemRpc_CreateAdmin_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _SystemRpc_UpdateAdmin_Handler,
		},
		{
			MethodName: "DeleteAdmin",
			Handler:    _SystemRpc_DeleteAdmin_Handler,
		},
		{
			MethodName: "RetrieveUserRole",
			Handler:    _SystemRpc_RetrieveUserRole_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _SystemRpc_UpdateUserRole_Handler,
		},
		{
			MethodName: "ContainsRole",
			Handler:    _SystemRpc_ContainsRole_Handler,
		},
		{
			MethodName: "RetrieveCommunity",
			Handler:    _SystemRpc_RetrieveCommunity_Handler,
		},
		{
			MethodName: "ListCommunity",
			Handler:    _SystemRpc_ListCommunity_Handler,
		},
		{
			MethodName: "CreateCommunity",
			Handler:    _SystemRpc_CreateCommunity_Handler,
		},
		{
			MethodName: "UpdateCommunity",
			Handler:    _SystemRpc_UpdateCommunity_Handler,
		},
		{
			MethodName: "DeleteCommunity",
			Handler:    _SystemRpc_DeleteCommunity_Handler,
		},
		{
			MethodName: "ListUseridByRole",
			Handler:    _SystemRpc_ListUseridByRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system.proto",
}
