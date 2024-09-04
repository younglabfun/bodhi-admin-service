// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package server

import (
	"context"

	"bodhiadmin/app/admin/rpc/internal/logic/userrole"
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"
)

type UserRoleServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedUserRoleServer
}

func NewUserRoleServer(svcCtx *svc.ServiceContext) *UserRoleServer {
	return &UserRoleServer{
		svcCtx: svcCtx,
	}
}

func (s *UserRoleServer) InsertUserRole(ctx context.Context, in *admin.UserRoleReq) (*admin.AffectedResp, error) {
	l := userrolelogic.NewInsertUserRoleLogic(ctx, s.svcCtx)
	return l.InsertUserRole(in)
}

func (s *UserRoleServer) UpdateUserRole(ctx context.Context, in *admin.UpdateUserRoleReq) (*admin.AffectedResp, error) {
	l := userrolelogic.NewUpdateUserRoleLogic(ctx, s.svcCtx)
	return l.UpdateUserRole(in)
}

func (s *UserRoleServer) RemoveUserRole(ctx context.Context, in *admin.Id) (*admin.AffectedResp, error) {
	l := userrolelogic.NewRemoveUserRoleLogic(ctx, s.svcCtx)
	return l.RemoveUserRole(in)
}

func (s *UserRoleServer) GetUserRoles(ctx context.Context, in *admin.Uuid) (*admin.UserRoleListResp, error) {
	l := userrolelogic.NewGetUserRolesLogic(ctx, s.svcCtx)
	return l.GetUserRoles(in)
}