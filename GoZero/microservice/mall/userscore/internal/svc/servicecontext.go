package svc

import (
	"userscore/database"
	"userscore/internal/config"
	"userscore/internal/dao"
	"userscore/internal/repo"
)

type ServiceContext struct {
	Config        config.Config
	UserScoreRepo repo.UserScoreRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserScoreRepo: dao.NewUserScoreDao(database.Connect(c.Mysql.DataSource, c.CacheRedis)),
	}
}
