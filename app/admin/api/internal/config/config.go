package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AdminRpcConf zrpc.RpcClientConf
	Auth         Auth
}
type Auth struct {
	AccessSecret  string
	AccessExpired int64
}
