package config

import "github.com/zeromicro/go-zero/zrpc"

// 这个config文件是对etc/user.yaml配置文件的一种映射
type Config struct {
	zrpc.RpcServerConf
	Mysql MysqlConfig
}

// Mysql的配置
type MysqlConfig struct {
	DataSource string
}
