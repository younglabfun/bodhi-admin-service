// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/logic/role"
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"
)

type RoleServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedRoleServer
}

func NewRoleServer(svcCtx *svc.ServiceContext) *RoleServer {
	return &RoleServer{
		svcCtx: svcCtx,
	}
}

func (s *RoleServer) InsertRole(ctx context.Context, in *admin.RoleReq) (*admin.AffectedResp, error) {
	l := rolelogic.NewInsertRoleLogic(ctx, s.svcCtx)
	return l.InsertRole(in)
}

func (s *RoleServer) UpdateRole(ctx context.Context, in *admin.RoleReq) (*admin.AffectedResp, error) {
	l := rolelogic.NewUpdateRoleLogic(ctx, s.svcCtx)
	return l.UpdateRole(in)
}

func (s *RoleServer) UpdateStatus(ctx context.Context, in *admin.UuidStatusReq) (*admin.AffectedResp, error) {
	l := rolelogic.NewUpdateStatusLogic(ctx, s.svcCtx)
	return l.UpdateStatus(in)
}

func (s *RoleServer) RemoveRole(ctx context.Context, in *admin.Uuid) (*admin.AffectedResp, error) {
	l := rolelogic.NewRemoveRoleLogic(ctx, s.svcCtx)
	return l.RemoveRole(in)
}

func (s *RoleServer) GetRole(ctx context.Context, in *admin.Uuid) (*admin.RoleUnit, error) {
	l := rolelogic.NewGetRoleLogic(ctx, s.svcCtx)
	return l.GetRole(in)
}

func (s *RoleServer) ListRole(ctx context.Context, in *admin.PageReq) (*admin.ListRoleResp, error) {
	l := rolelogic.NewListRoleLogic(ctx, s.svcCtx)
	return l.ListRole(in)
}

func (s *RoleServer) GetList(ctx context.Context, in *admin.Empty) (*admin.RoleListResp, error) {
	l := rolelogic.NewGetListLogic(ctx, s.svcCtx)
	return l.GetList(in)
}
