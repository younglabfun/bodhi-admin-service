package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"os"

	"bodhiadmin/app/admin/api/internal/config"
	"bodhiadmin/app/admin/api/internal/handler"
	"bodhiadmin/app/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.SetUp(c.Log)
	// 输出至屏幕
	if c.AdminConf.Debug {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	logx.Infof("------ Start Server %s %s", c.AdminConf.App, c.AdminConf.Version)
	logx.Infof("-- Init Log path: %s, level: %s, debug: %v", c.Log.Path, c.Log.Level, c.AdminConf.Debug)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
