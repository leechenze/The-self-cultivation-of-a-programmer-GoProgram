package svc

import (
	"user/database"
	"user/internal/config"
	"user/internal/dao"
	"user/internal/repo"
)

// svc 就是service层，依赖的资源池，一些数据库连接通过这里导入调用，
type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(database.Connect(c.Mysql.DataSource)), // userDao实现了userRepo的save方法，所以说userDao是userRepo的实现类
	}
}
