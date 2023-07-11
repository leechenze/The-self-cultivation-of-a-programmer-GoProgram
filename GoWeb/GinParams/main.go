package GinParams

import (
	"github.com/gin-gonic/gin"
)

func Main() {
	/** 定义Gin路由引擎 */
	engine := gin.Default()

	/** get参数请求 */
	// engine.GET("/", GetParams)

	/** form参数请求 */
	// engine.LoadHTMLFiles("./GinParams/template/login.tmpl", "./GinParams/template/login-after.tmpl")
	// // 返回 Login 页面
	// engine.GET("/login", PostParamsBefore)
	// // 登录逻辑 并且返回 Home 页面
	// engine.POST("/login", PostParamsAfter)

	/** 获取 Uri 路径参数 */
	// engine.GET("/:name/:age", UriParams)

	/** 参数绑定 */
	engine.GET("/", BindParams)

	/** 运行Gin程序 */
	engine.Run(":9090")
}
