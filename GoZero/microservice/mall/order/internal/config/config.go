package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 这个配置映射的就是 etc下的配置文件：order-api.yaml
type Config struct {
	rest.RestConf
	// 去etcd获取user.rpc的地址
	UserRPC zrpc.RpcClientConf
}
