package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	AdminConf    AdminConf
	AdminRpcConf zrpc.RpcClientConf
	Auth         Auth
}
type Auth struct {
	AccessSecret  string
	AccessExpired int64
}
type AdminConf struct {
	App     string `json:"app"`
	Version string `json:"version"`
	Debug   bool   `json:"debug"`
}
