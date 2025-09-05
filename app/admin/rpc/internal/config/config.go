package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AdminConf AdminConf
	AuthConf  Auth
	MySql     MySql
	RedisConf Redis `json:"optional"`
}

type MySql struct {
	Database string
	Host     string
	Port     int64
	User     string
	Password string
}

type Redis struct {
	Host     string `json:",optional"`
	Port     int64  `json:",optional"`
	DBIndex  int    `json:",optional"`
	Password string `json:",optional"`
}

type AdminConf struct {
	App            string `json:"app"`
	Version        string `json:"version"`
	Debug          bool   `json:"debug"`
	Salt           string `json:"salt"`
	Master         string `json:"master"`
	RefreshExpired int64  `json:"refreshExpired"`
}

type Auth struct {
	AccessSecret  string
	AccessExpired int64
}
