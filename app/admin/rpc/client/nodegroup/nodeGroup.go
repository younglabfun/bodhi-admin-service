// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package nodegroup

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

	NodeGroup interface {
		InsertNodeGroup(ctx context.Context, in *NodeGroupReq, opts ...grpc.CallOption) (*AffectedResp, error)
		UpdateNodeGroup(ctx context.Context, in *NodeGroupReq, opts ...grpc.CallOption) (*AffectedResp, error)
		RemoveNodeGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*AffectedResp, error)
		GetNodeGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*NodeGroupUnit, error)
		GetNodeGroupList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeGroupList, error)
	}

	defaultNodeGroup struct {
		cli zrpc.Client
	}
)

func NewNodeGroup(cli zrpc.Client) NodeGroup {
	return &defaultNodeGroup{
		cli: cli,
	}
}

func (m *defaultNodeGroup) InsertNodeGroup(ctx context.Context, in *NodeGroupReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewNodeGroupClient(m.cli.Conn())
	return client.InsertNodeGroup(ctx, in, opts...)
}

func (m *defaultNodeGroup) UpdateNodeGroup(ctx context.Context, in *NodeGroupReq, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewNodeGroupClient(m.cli.Conn())
	return client.UpdateNodeGroup(ctx, in, opts...)
}

func (m *defaultNodeGroup) RemoveNodeGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*AffectedResp, error) {
	client := admin.NewNodeGroupClient(m.cli.Conn())
	return client.RemoveNodeGroup(ctx, in, opts...)
}

func (m *defaultNodeGroup) GetNodeGroup(ctx context.Context, in *Id, opts ...grpc.CallOption) (*NodeGroupUnit, error) {
	client := admin.NewNodeGroupClient(m.cli.Conn())
	return client.GetNodeGroup(ctx, in, opts...)
}

func (m *defaultNodeGroup) GetNodeGroupList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NodeGroupList, error) {
	client := admin.NewNodeGroupClient(m.cli.Conn())
	return client.GetNodeGroupList(ctx, in, opts...)
}