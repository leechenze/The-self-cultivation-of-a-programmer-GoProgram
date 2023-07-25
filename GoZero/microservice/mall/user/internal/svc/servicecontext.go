package svc

import (
	"user/internal/config"
	"user/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
