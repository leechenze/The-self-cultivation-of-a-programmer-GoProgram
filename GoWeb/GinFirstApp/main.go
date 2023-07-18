package GinFirstApp

import "github.com/gin-gonic/gin"

func sayHello(con *gin.Context) {
	con.JSON(200, gin.H{
		"message": "hello gin",
	})
}

func Main() {

	// 返回默认的路由引擎
	engine := gin.Default()

	// 请求方式及执行逻辑
	engine.GET("/hello", sayHello)

	// 启动服务
	engine.Run(":9090")
}
