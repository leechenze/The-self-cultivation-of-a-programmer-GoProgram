package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	ConnCache sqlc.CachedConn
}

func Connect(datasource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(datasource)
	cacheConn := sqlc.NewConn(sqlConn, conf)
	databaseConn := &DBConn{
		Conn: sqlConn,
	}
	// 如果conf参数有值时才用redis缓存连接（cacheConn）
	if conf != nil {
		databaseConn.ConnCache = cacheConn
	}

	return databaseConn
}
