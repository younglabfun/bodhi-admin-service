package main

import (
	"flag"
	"fmt"

	"bodhiadmin/app/admin/rpc/internal/config"
	accountServer "bodhiadmin/app/admin/rpc/internal/server/account"
	menuServer "bodhiadmin/app/admin/rpc/internal/server/menu"
	nodeServer "bodhiadmin/app/admin/rpc/internal/server/node"
	nodegroupServer "bodhiadmin/app/admin/rpc/internal/server/nodegroup"
	roleServer "bodhiadmin/app/admin/rpc/internal/server/role"
	userServer "bodhiadmin/app/admin/rpc/internal/server/user"
	userroleServer "bodhiadmin/app/admin/rpc/internal/server/userrole"
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAccountServer(grpcServer, accountServer.NewAccountServer(ctx))
		admin.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))
		admin.RegisterNodeServer(grpcServer, nodeServer.NewNodeServer(ctx))
		admin.RegisterNodeGroupServer(grpcServer, nodegroupServer.NewNodeGroupServer(ctx))
		admin.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		admin.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		admin.RegisterUserRoleServer(grpcServer, userroleServer.NewUserRoleServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
