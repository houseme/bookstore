package svc

import (
	"github.com/houseme/bookstore/api/internal/config"
	"github.com/houseme/bookstore/rpc/add/adder"
	"github.com/houseme/bookstore/rpc/check/checker"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder     // 手动代码
	Checker checker.Checker // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),       // 手动代码
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)), // 手动代码
	}
}
