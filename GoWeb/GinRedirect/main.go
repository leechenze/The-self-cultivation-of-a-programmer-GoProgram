package GinRedirect

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {

	engine := gin.Default()

	/** 重定向方式一 */
	engine.GET("/index", func(context *gin.Context) {
		// 路由重定向（StatusMovedPermanently，谨慎使用将永久重定向地址到 google.com）
		// context.Redirect(http.StatusMovedPermanently, "https://www.google.com")
		// 路由重定向（StatusTemporaryRedirect，是临时重定向）
		context.Redirect(http.StatusTemporaryRedirect, "https://www.baidu.com")
	})

	/** 重定向方式二（路由转发） */
	engine.GET("/users", func(context *gin.Context) {
		// 强制修改请求的URI
		context.Request.URL.Path = "/orders"
		// 继续后续的处理
		engine.HandleContext(context)
	})
	engine.GET("/orders", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "redirect to orders page",
		})
	})

	engine.Run(":9090")

}
