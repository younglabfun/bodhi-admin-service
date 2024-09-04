package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource string
	}
	JwtAuthConf JwtAuthConf
	AdminConf   AdminConf
}

type AdminConf struct {
	Salt           string `json:"salt"`
	Master         string `json:"master"`
	RefreshExpired int64  `json:"refreshExpired"`
}

type JwtAuthConf struct {
	Secret  string `json:"secret"`
	Expired int64  `json:"expired"`
}
