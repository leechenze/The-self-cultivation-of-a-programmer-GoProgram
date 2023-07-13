package GinMiddleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(ctx *gin.Context) {
	println("indexHandler is run")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件
func middleWareFn(ctx *gin.Context) {
	println("middleware func in ...")
	// 计时程序运行耗时
	startTime := time.Now()
	ctx.Next() // next就是继续完后执行处理函数（indexHandler）
	// ctx.Abort() // abort就是阻止后续处理函数的执行（indexHandler）
	// Since是从startTime 到 当前时间的中间时差
	sinceTime := time.Since(startTime)
	fmt.Printf("cost time is: %v\n", sinceTime)
	println("middleware func out ...")
}

// 更为严谨的中间件写法示例，比如写一个权限登录的中间件。
func authMiddleWareFn() gin.HandlerFunc {
	// 连接数据库等一些其他准备工作
	return func(ctx *gin.Context) {
		// 中间件逻辑编写
		// Set方法可以给当前上下文内写入一些值，这样在路由处理函数中的通过context共享。
		ctx.Set("isLogin", "0")
	}
}

func Main() {
	engine := gin.Default()

	// 方式一：给单个路由添加中间件（给index路由之前添加中间执行函数：middlewareFn）
	engine.GET("/index", middleWareFn, indexHandler)

	// 方式二：给全局添加中间件（相当于给所有路由之前都添加中间执行函数：middlewareFn）
	engine.Use(middleWareFn, authMiddleWareFn())
	engine.GET("/shop", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "shop"})
	})
	engine.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "users"})
	})
	engine.GET("/orders", func(ctx *gin.Context) {
		isLogin, _ := ctx.Get("isLogin")
		ctx.JSON(http.StatusOK, gin.H{"msg": "orders", "isLogin": isLogin})
	})

	engine.Run(":9090")
}
