package GinRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {

	engine := gin.Default()

	// Any 可以接收所有的请求方式：GET, POST, DELETE...
	engine.Any("/method", func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{
				"requestMethod": http.MethodPost,
			})
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{
				"requestMethod": http.MethodGet,
			})
		case http.MethodDelete:
			context.JSON(http.StatusOK, gin.H{
				"requestMethod": http.MethodDelete,
			})
		}
	})

	// NoRoute 可以接收所有未知路由
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"message": "空路由：NoRoute",
		})
	})

	// 路由组
	// 提取公共前缀，创建一个路由组
	videoGroup := engine.Group("/video")
	// 这个花括号不是必须的,只是为了代码整洁.
	{
		videoGroup.GET("/aaa", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "video前缀的aaa接口数据"})
		})
		videoGroup.GET("/bbb", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "video前缀的bbb接口数据"})
		})
		videoGroup.GET("/ccc", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "video前缀的ccc接口数据"})
		})
		// 路由组嵌套
		videoDddGroup := videoGroup.Group("/ddd")
		{
			videoDddGroup.GET("aaa", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"msg": "video/ddd前缀的aaa接口数据"})
			})
		}
	}

	/** 启动路由 */
	engine.Run(":9090")

}
