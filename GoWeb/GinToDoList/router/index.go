package router

import (
	"GoWeb/GinToDoList/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	// 创建Gin路由
	engine := gin.Default()
	// <link href=/static/css/app.8eeeaf31.css rel=preload as=style>
	// 就是指定如上 href的地址的 /static 目录对应目录是 ./GinToDoList/static
	engine.Static("/static", "./GinToDoList/static")
	// 加载模版文件
	engine.LoadHTMLGlob("./GinToDoList/templates/*")
	// 建立 / 路由请求
	engine.GET("/", controller.IndexHandler)

	// 请求接口分组
	v1Group := engine.Group("v1")
	{
		/** 待办事项 */
		// 添加
		v1Group.POST("/todo", controller.AddToDo)
		// 查询所有
		v1Group.GET("/todo", controller.QueryToDoList)
		// 查询一个
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {})
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateToDo)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteToDo)
	}
	return engine
}
