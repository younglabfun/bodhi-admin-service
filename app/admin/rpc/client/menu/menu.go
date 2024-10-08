// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package menu

import (
	"context"

	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AffectedResp      = admin.AffectedResp
	BatchIdsReq       = admin.BatchIdsReq
	Empty             = admin.Empty
	Id                = admin.Id
	ListMenuReq       = admin.ListMenuReq
	ListNodeResp      = admin.ListNodeResp
	ListRoleResp      = admin.ListRoleResp
	ListUserResp      = admin.ListUserResp
	LoginReq          = admin.LoginReq
	LoginResp         = admin.LoginResp
	MenuListResp      = admin.MenuListResp
	MenuReq           = admin.MenuReq
	MenuResp          = admin.MenuResp
	MenuTypeResp      = admin.MenuTypeResp
	MenuUnit          = admin.MenuUnit
	MoveReq           = admin.MoveReq
	NewUserReq        = admin.NewUserReq
	NodeGroupList     = admin.NodeGroupList
	NodeGroupReq      = admin.NodeGroupReq
	NodeGroupUnit     = admin.NodeGroupUnit
	NodeListResp      = admin.NodeListResp
	NodeReq           = admin.NodeReq
	NodeUnit          = admin.NodeUnit
	PageReq           = admin.PageReq
	PasswordReq       = admin.PasswordReq
	PermissionResp    = admin.PermissionResp
	RegisterReq       = admin.RegisterReq
	RoleListResp      = admin.RoleListResp
	RoleReq           = admin.RoleReq
	RoleUnit          = admin.RoleUnit
	StatusReq         = admin.StatusReq
	TokenReq          = admin.TokenReq
	TokenResp         = admin.TokenResp
	TokenUnit         = admin.TokenUnit
	UpdateUserRoleReq = admin.UpdateUserRoleReq
	UserPasswordReq   = admin.UserPasswordReq
	UserReq           = admin.UserReq
	UserResp          = admin.UserResp
	UserRoleListResp  = admin.UserRoleListResp
	UserRoleReq       = admin.UserRoleReq
	UserRoleUnit      = admin.UserRoleUnit
	UserUnit          = admin.UserUnit
	Uuid              = admin.Uuid
	UuidStatusReq     = admin.UuidStatusReq

	Menu interface {
		InsertMenu(ctx context.Context, in *MenuReq, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdateMenu(ctx context.Context, in *MenuReq, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdateStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*AffectedResp, error)
		RemoveMenu(ctx context.Context, in *Id, opts ...grpc.CallOption) (*AffectedResp, error)
		GetMenu(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MenuResp, error)
		GetMenuListByPid(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*MenuListResp, error)
		GetMenuListByType(ctx context.Context, in *MenuTypeResp, opts ...grpc.CallOption) (*MenuListResp, error)
	}

	defaultMenu struct {
		cli zrpc.Client
	}
)

func NewMenu(cli zrpc.Client) Menu {
	return &defaultMenu{
		cli: cli,
	}
}

func (m *defaultMenu) InsertMenu(ctx context.Context, in *MenuReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.InsertMenu(ctx, in, opts...)
}

func (m *defaultMenu) UpdateMenu(ctx context.Context, in *MenuReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.UpdateMenu(ctx, in, opts...)
}

func (m *defaultMenu) UpdateStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}

func (m *defaultMenu) RemoveMenu(ctx context.Context, in *Id, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.RemoveMenu(ctx, in, opts...)
}

func (m *defaultMenu) GetMenu(ctx context.Context, in *Id, opts ...grpc.CallOption) (*MenuResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.GetMenu(ctx, in, opts...)
}

func (m *defaultMenu) GetMenuListByPid(ctx context.Context, in *ListMenuReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.GetMenuListByPid(ctx, in, opts...)
}

func (m *defaultMenu) GetMenuListByType(ctx context.Context, in *MenuTypeResp, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := admin.NewMenuClient(m.cli.Conn())
	return client.GetMenuListByType(ctx, in, opts...)
}
