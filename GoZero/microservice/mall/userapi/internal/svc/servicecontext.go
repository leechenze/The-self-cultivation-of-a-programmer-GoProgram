package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"rpc-common/userclient"
	"userapi/internal/config"
	"userapi/internal/middlewares"
)

// svc 就是service层，其中是一些依赖的资源池
type ServiceContext struct {
	Config         config.Config
	UserRPC        userclient.User
	UserMiddleware *middlewares.UserMiddleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserRPC:        userclient.NewUser(zrpc.MustNewClient(c.UserRPC)),
		UserMiddleware: middlewares.NewUserMiddleware(),
	}
}
