// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package systemrpc

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-system-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Admin                 = pb2.Admin
	Community             = pb2.Community
	ContainsRoleReq       = pb2.ContainsRoleReq
	ContainsRoleResp      = pb2.ContainsRoleResp
	CreateAdminReq        = pb2.CreateAdminReq
	CreateAdminResp       = pb2.CreateAdminResp
	CreateCommunityReq    = pb2.CreateCommunityReq
	CreateCommunityResp   = pb2.CreateCommunityResp
	CreateNewsReq         = pb2.CreateNewsReq
	CreateNewsResp        = pb2.CreateNewsResp
	CreateNoticeReq       = pb2.CreateNoticeReq
	CreateNoticeResp      = pb2.CreateNoticeResp
	DeleteAdminReq        = pb2.DeleteAdminReq
	DeleteAdminResp       = pb2.DeleteAdminResp
	DeleteCommunityReq    = pb2.DeleteCommunityReq
	DeleteCommunityResp   = pb2.DeleteCommunityResp
	DeleteNewsReq         = pb2.DeleteNewsReq
	DeleteNewsResp        = pb2.DeleteNewsResp
	DeleteNoticeReq       = pb2.DeleteNoticeReq
	DeleteNoticeResp      = pb2.DeleteNoticeResp
	ListAdminReq          = pb2.ListAdminReq
	ListAdminResp         = pb2.ListAdminResp
	ListCommunityReq      = pb2.ListCommunityReq
	ListCommunityResp     = pb2.ListCommunityResp
	ListNewsReq           = pb2.ListNewsReq
	ListNewsResp          = pb2.ListNewsResp
	ListNoticeReq         = pb2.ListNoticeReq
	ListNoticeResp        = pb2.ListNoticeResp
	ListUseridReq         = pb2.ListUseridReq
	ListUseridResp        = pb2.ListUseridResp
	News                  = pb2.News
	Notice                = pb2.Notice
	RetrieveAdminReq      = pb2.RetrieveAdminReq
	RetrieveAdminResp     = pb2.RetrieveAdminResp
	RetrieveCommunityReq  = pb2.RetrieveCommunityReq
	RetrieveCommunityResp = pb2.RetrieveCommunityResp
	RetrieveNewsReq       = pb2.RetrieveNewsReq
	RetrieveNewsResp      = pb2.RetrieveNewsResp
	RetrieveNoticeReq     = pb2.RetrieveNoticeReq
	RetrieveNoticeResp    = pb2.RetrieveNoticeResp
	RetrieveUserRoleReq   = pb2.RetrieveUserRoleReq
	RetrieveUserRoleResp  = pb2.RetrieveUserRoleResp
	Role                  = pb2.Role
	UpdateAdminReq        = pb2.UpdateAdminReq
	UpdateAdminResp       = pb2.UpdateAdminResp
	UpdateCommunityReq    = pb2.UpdateCommunityReq
	UpdateCommunityResp   = pb2.UpdateCommunityResp
	UpdateNewsReq         = pb2.UpdateNewsReq
	UpdateNewsResp        = pb2.UpdateNewsResp
	UpdateNoticeReq       = pb2.UpdateNoticeReq
	UpdateNoticeResp      = pb2.UpdateNoticeResp
	UpdateUserRoleReq     = pb2.UpdateUserRoleReq
	UpdateUserRoleResp    = pb2.UpdateUserRoleResp

	SystemRpc interface {
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

	defaultSystemRpc struct {
		cli zrpc.Client
	}
)

func NewSystemRpc(cli zrpc.Client) SystemRpc {
	return &defaultSystemRpc{
		cli: cli,
	}
}

func (m *defaultSystemRpc) RetrieveNotice(ctx context.Context, in *RetrieveNoticeReq, opts ...grpc.CallOption) (*RetrieveNoticeResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.RetrieveNotice(ctx, in, opts...)
}

func (m *defaultSystemRpc) ListNotice(ctx context.Context, in *ListNoticeReq, opts ...grpc.CallOption) (*ListNoticeResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ListNotice(ctx, in, opts...)
}

func (m *defaultSystemRpc) CreateNotice(ctx context.Context, in *CreateNoticeReq, opts ...grpc.CallOption) (*CreateNoticeResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.CreateNotice(ctx, in, opts...)
}

func (m *defaultSystemRpc) UpdateNotice(ctx context.Context, in *UpdateNoticeReq, opts ...grpc.CallOption) (*UpdateNoticeResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.UpdateNotice(ctx, in, opts...)
}

func (m *defaultSystemRpc) DeleteNotice(ctx context.Context, in *DeleteNoticeReq, opts ...grpc.CallOption) (*DeleteNoticeResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.DeleteNotice(ctx, in, opts...)
}

func (m *defaultSystemRpc) RetrieveNews(ctx context.Context, in *RetrieveNewsReq, opts ...grpc.CallOption) (*RetrieveNewsResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.RetrieveNews(ctx, in, opts...)
}

func (m *defaultSystemRpc) ListNews(ctx context.Context, in *ListNewsReq, opts ...grpc.CallOption) (*ListNewsResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ListNews(ctx, in, opts...)
}

func (m *defaultSystemRpc) CreateNews(ctx context.Context, in *CreateNewsReq, opts ...grpc.CallOption) (*CreateNewsResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.CreateNews(ctx, in, opts...)
}

func (m *defaultSystemRpc) UpdateNews(ctx context.Context, in *UpdateNewsReq, opts ...grpc.CallOption) (*UpdateNewsResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.UpdateNews(ctx, in, opts...)
}

func (m *defaultSystemRpc) DeleteNews(ctx context.Context, in *DeleteNewsReq, opts ...grpc.CallOption) (*DeleteNewsResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.DeleteNews(ctx, in, opts...)
}

func (m *defaultSystemRpc) RetrieveAdmin(ctx context.Context, in *RetrieveAdminReq, opts ...grpc.CallOption) (*RetrieveAdminResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.RetrieveAdmin(ctx, in, opts...)
}

func (m *defaultSystemRpc) ListAdmin(ctx context.Context, in *ListAdminReq, opts ...grpc.CallOption) (*ListAdminResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ListAdmin(ctx, in, opts...)
}

func (m *defaultSystemRpc) CreateAdmin(ctx context.Context, in *CreateAdminReq, opts ...grpc.CallOption) (*CreateAdminResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.CreateAdmin(ctx, in, opts...)
}

func (m *defaultSystemRpc) UpdateAdmin(ctx context.Context, in *UpdateAdminReq, opts ...grpc.CallOption) (*UpdateAdminResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.UpdateAdmin(ctx, in, opts...)
}

func (m *defaultSystemRpc) DeleteAdmin(ctx context.Context, in *DeleteAdminReq, opts ...grpc.CallOption) (*DeleteAdminResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.DeleteAdmin(ctx, in, opts...)
}

// 获取用户的所有角色
func (m *defaultSystemRpc) RetrieveUserRole(ctx context.Context, in *RetrieveUserRoleReq, opts ...grpc.CallOption) (*RetrieveUserRoleResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.RetrieveUserRole(ctx, in, opts...)
}

// 更新用户的角色
func (m *defaultSystemRpc) UpdateUserRole(ctx context.Context, in *UpdateUserRoleReq, opts ...grpc.CallOption) (*UpdateUserRoleResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.UpdateUserRole(ctx, in, opts...)
}

func (m *defaultSystemRpc) ContainsRole(ctx context.Context, in *ContainsRoleReq, opts ...grpc.CallOption) (*ContainsRoleResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ContainsRole(ctx, in, opts...)
}

func (m *defaultSystemRpc) RetrieveCommunity(ctx context.Context, in *RetrieveCommunityReq, opts ...grpc.CallOption) (*RetrieveCommunityResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.RetrieveCommunity(ctx, in, opts...)
}

func (m *defaultSystemRpc) ListCommunity(ctx context.Context, in *ListCommunityReq, opts ...grpc.CallOption) (*ListCommunityResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ListCommunity(ctx, in, opts...)
}

func (m *defaultSystemRpc) CreateCommunity(ctx context.Context, in *CreateCommunityReq, opts ...grpc.CallOption) (*CreateCommunityResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.CreateCommunity(ctx, in, opts...)
}

func (m *defaultSystemRpc) UpdateCommunity(ctx context.Context, in *UpdateCommunityReq, opts ...grpc.CallOption) (*UpdateCommunityResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.UpdateCommunity(ctx, in, opts...)
}

func (m *defaultSystemRpc) DeleteCommunity(ctx context.Context, in *DeleteCommunityReq, opts ...grpc.CallOption) (*DeleteCommunityResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.DeleteCommunity(ctx, in, opts...)
}

func (m *defaultSystemRpc) ListUseridByRole(ctx context.Context, in *ListUseridReq, opts ...grpc.CallOption) (*ListUseridResp, error) {
	client := pb2.NewSystemRpcClient(m.cli.Conn())
	return client.ListUseridByRole(ctx, in, opts...)
}
