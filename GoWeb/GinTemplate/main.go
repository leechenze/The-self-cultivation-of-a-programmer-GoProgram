package GinTemplate

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func Main() {
	/** 定义Gin路由（引擎）*/
	engine := gin.Default()
	/** 静态文件处理 */
	// 第一个参数是相对路径，在模版中使用：static/index.css
	engine.Static("/static", "./GinTemplate/static")
	/** 自定义模版函数 */
	engine.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	/** 模版解析 */
	// 加载一个模版
	// engine.LoadHTMLFiles("./GinTemplate/template/index.tmpl")
	// 加载所有模版
	engine.LoadHTMLGlob("./GinTemplate/template/**/*.tmpl")
	/** 模版渲染 */
	engine.GET("posts/index", func(context *gin.Context) {
		// H 相当于Gin定义的一个便捷的Map
		context.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title":       "posts/index.tmpl",
			"siteAddress": "<a href='http://www.leechenze.com'>李晨泽的博客<a/>",
		})
	})
	engine.GET("users/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title":       "users/index.tmpl",
			"siteAddress": "<a href='http://www.leechenze.com'>李晨泽的博客<a/>",
		})
	})
	/** 运行Gin程序 */
	engine.Run(":9090")
}
