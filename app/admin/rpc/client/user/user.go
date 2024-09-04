// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package user

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

	User interface {
		InsertUser(ctx context.Context, in *NewUserReq, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdateStatus(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdatePassword(ctx context.Context, in *UserPasswordReq, opts ...grpc.CallOption) (*AffectedResp, error)
		RemoveUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AffectedResp, error)
		GetUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*UserUnit, error)
		ListUser(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*ListUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) InsertUser(ctx context.Context, in *NewUserReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.InsertUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateUser(ctx context.Context, in *UserReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUser) UpdateStatus(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.UpdateStatus(ctx, in, opts...)
}

func (m *defaultUser) UpdatePassword(ctx context.Context, in *UserPasswordReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.UpdatePassword(ctx, in, opts...)
}

func (m *defaultUser) RemoveUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.RemoveUser(ctx, in, opts...)
}

func (m *defaultUser) GetUser(ctx context.Context, in *Uuid, opts ...grpc.CallOption) (*UserUnit, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

func (m *defaultUser) ListUser(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*ListUserResp, error) {
	client := admin.NewUserClient(m.cli.Conn())
	return client.ListUser(ctx, in, opts...)
}