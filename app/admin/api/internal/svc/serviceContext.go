package svc

import (
	"bodhiadmin/app/admin/api/internal/config"
	"bodhiadmin/app/admin/rpc/client/account"
	"bodhiadmin/app/admin/rpc/client/menu"
	"bodhiadmin/app/admin/rpc/client/node"
	"bodhiadmin/app/admin/rpc/client/nodegroup"
	"bodhiadmin/app/admin/rpc/client/role"
	"bodhiadmin/app/admin/rpc/client/user"
	"bodhiadmin/app/admin/rpc/client/userrole"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AccountRpc   account.Account
	MenuRpc      menu.Menu
	NodeRpc      node.Node
	NodeGroupRpc nodegroup.NodeGroup
	RoleRpc      role.Role
	UserRpc      user.User
	UserRoleRpc  userrole.UserRole
}

func NewServiceContext(c config.Config) *ServiceContext {
	accountRpc := account.NewAccount(zrpc.MustNewClient(c.AdminRpcConf))
	menuRpc := menu.NewMenu(zrpc.MustNewClient(c.AdminRpcConf))
	nodeRpc := node.NewNode(zrpc.MustNewClient(c.AdminRpcConf))
	nodeGroupRpc := nodegroup.NewNodeGroup(zrpc.MustNewClient(c.AdminRpcConf))
	roleRpc := role.NewRole(zrpc.MustNewClient(c.AdminRpcConf))
	userRpc := user.NewUser(zrpc.MustNewClient(c.AdminRpcConf))
	userRoleRpc := userrole.NewUserRole(zrpc.MustNewClient(c.AdminRpcConf))

	return &ServiceContext{
		Config: c,

		AccountRpc:   accountRpc,
		MenuRpc:      menuRpc,
		NodeRpc:      nodeRpc,
		NodeGroupRpc: nodeGroupRpc,
		RoleRpc:      roleRpc,
		UserRpc:      userRpc,
		UserRoleRpc:  userRoleRpc,
	}
}
