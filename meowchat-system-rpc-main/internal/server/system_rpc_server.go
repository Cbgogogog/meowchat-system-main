// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package server

import (
	"context"

	"github.com/xh-polaris/meowchat-system-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-system-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-system-rpc/pb"
)

type SystemRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSystemRpcServer
}

func NewSystemRpcServer(svcCtx *svc.ServiceContext) *SystemRpcServer {
	return &SystemRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *SystemRpcServer) RetrieveNotice(ctx context.Context, in *pb.RetrieveNoticeReq) (*pb.RetrieveNoticeResp, error) {
	l := logic.NewRetrieveNoticeLogic(ctx, s.svcCtx)
	return l.RetrieveNotice(in)
}

func (s *SystemRpcServer) ListNotice(ctx context.Context, in *pb.ListNoticeReq) (*pb.ListNoticeResp, error) {
	l := logic.NewListNoticeLogic(ctx, s.svcCtx)
	return l.ListNotice(in)
}

func (s *SystemRpcServer) CreateNotice(ctx context.Context, in *pb.CreateNoticeReq) (*pb.CreateNoticeResp, error) {
	l := logic.NewCreateNoticeLogic(ctx, s.svcCtx)
	return l.CreateNotice(in)
}

func (s *SystemRpcServer) UpdateNotice(ctx context.Context, in *pb.UpdateNoticeReq) (*pb.UpdateNoticeResp, error) {
	l := logic.NewUpdateNoticeLogic(ctx, s.svcCtx)
	return l.UpdateNotice(in)
}

func (s *SystemRpcServer) DeleteNotice(ctx context.Context, in *pb.DeleteNoticeReq) (*pb.DeleteNoticeResp, error) {
	l := logic.NewDeleteNoticeLogic(ctx, s.svcCtx)
	return l.DeleteNotice(in)
}

func (s *SystemRpcServer) RetrieveNews(ctx context.Context, in *pb.RetrieveNewsReq) (*pb.RetrieveNewsResp, error) {
	l := logic.NewRetrieveNewsLogic(ctx, s.svcCtx)
	return l.RetrieveNews(in)
}

func (s *SystemRpcServer) ListNews(ctx context.Context, in *pb.ListNewsReq) (*pb.ListNewsResp, error) {
	l := logic.NewListNewsLogic(ctx, s.svcCtx)
	return l.ListNews(in)
}

func (s *SystemRpcServer) CreateNews(ctx context.Context, in *pb.CreateNewsReq) (*pb.CreateNewsResp, error) {
	l := logic.NewCreateNewsLogic(ctx, s.svcCtx)
	return l.CreateNews(in)
}

func (s *SystemRpcServer) UpdateNews(ctx context.Context, in *pb.UpdateNewsReq) (*pb.UpdateNewsResp, error) {
	l := logic.NewUpdateNewsLogic(ctx, s.svcCtx)
	return l.UpdateNews(in)
}

func (s *SystemRpcServer) DeleteNews(ctx context.Context, in *pb.DeleteNewsReq) (*pb.DeleteNewsResp, error) {
	l := logic.NewDeleteNewsLogic(ctx, s.svcCtx)
	return l.DeleteNews(in)
}

func (s *SystemRpcServer) RetrieveAdmin(ctx context.Context, in *pb.RetrieveAdminReq) (*pb.RetrieveAdminResp, error) {
	l := logic.NewRetrieveAdminLogic(ctx, s.svcCtx)
	return l.RetrieveAdmin(in)
}

func (s *SystemRpcServer) ListAdmin(ctx context.Context, in *pb.ListAdminReq) (*pb.ListAdminResp, error) {
	l := logic.NewListAdminLogic(ctx, s.svcCtx)
	return l.ListAdmin(in)
}

func (s *SystemRpcServer) CreateAdmin(ctx context.Context, in *pb.CreateAdminReq) (*pb.CreateAdminResp, error) {
	l := logic.NewCreateAdminLogic(ctx, s.svcCtx)
	return l.CreateAdmin(in)
}

func (s *SystemRpcServer) UpdateAdmin(ctx context.Context, in *pb.UpdateAdminReq) (*pb.UpdateAdminResp, error) {
	l := logic.NewUpdateAdminLogic(ctx, s.svcCtx)
	return l.UpdateAdmin(in)
}

func (s *SystemRpcServer) DeleteAdmin(ctx context.Context, in *pb.DeleteAdminReq) (*pb.DeleteAdminResp, error) {
	l := logic.NewDeleteAdminLogic(ctx, s.svcCtx)
	return l.DeleteAdmin(in)
}

// 获取用户的所有角色
func (s *SystemRpcServer) RetrieveUserRole(ctx context.Context, in *pb.RetrieveUserRoleReq) (*pb.RetrieveUserRoleResp, error) {
	l := logic.NewRetrieveUserRoleLogic(ctx, s.svcCtx)
	return l.RetrieveUserRole(in)
}

// 更新用户的角色
func (s *SystemRpcServer) UpdateUserRole(ctx context.Context, in *pb.UpdateUserRoleReq) (*pb.UpdateUserRoleResp, error) {
	l := logic.NewUpdateUserRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserRole(in)
}

func (s *SystemRpcServer) ContainsRole(ctx context.Context, in *pb.ContainsRoleReq) (*pb.ContainsRoleResp, error) {
	l := logic.NewContainsRoleLogic(ctx, s.svcCtx)
	return l.ContainsRole(in)
}

func (s *SystemRpcServer) RetrieveCommunity(ctx context.Context, in *pb.RetrieveCommunityReq) (*pb.RetrieveCommunityResp, error) {
	l := logic.NewRetrieveCommunityLogic(ctx, s.svcCtx)
	return l.RetrieveCommunity(in)
}

func (s *SystemRpcServer) ListCommunity(ctx context.Context, in *pb.ListCommunityReq) (*pb.ListCommunityResp, error) {
	l := logic.NewListCommunityLogic(ctx, s.svcCtx)
	return l.ListCommunity(in)
}

func (s *SystemRpcServer) CreateCommunity(ctx context.Context, in *pb.CreateCommunityReq) (*pb.CreateCommunityResp, error) {
	l := logic.NewCreateCommunityLogic(ctx, s.svcCtx)
	return l.CreateCommunity(in)
}

func (s *SystemRpcServer) UpdateCommunity(ctx context.Context, in *pb.UpdateCommunityReq) (*pb.UpdateCommunityResp, error) {
	l := logic.NewUpdateCommunityLogic(ctx, s.svcCtx)
	return l.UpdateCommunity(in)
}

func (s *SystemRpcServer) DeleteCommunity(ctx context.Context, in *pb.DeleteCommunityReq) (*pb.DeleteCommunityResp, error) {
	l := logic.NewDeleteCommunityLogic(ctx, s.svcCtx)
	return l.DeleteCommunity(in)
}

func (s *SystemRpcServer) ListUseridByRole(ctx context.Context, in *pb.ListUseridReq) (*pb.ListUseridResp, error) {
	l := logic.NewListUseridByRoleLogic(ctx, s.svcCtx)
	return l.ListUseridByRole(in)
}