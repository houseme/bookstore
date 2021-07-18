package config

import (
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Add   zrpc.RpcClientConf // 手动代码
	Check zrpc.RpcClientConf // 手动代码
}
