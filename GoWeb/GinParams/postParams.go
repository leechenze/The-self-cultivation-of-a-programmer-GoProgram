package GinParams

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostParamsBefore(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", nil)

}

func PostParamsAfter(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	// 和 Query 同样还有 DefaultPostForm 和 GetPostForm，这里就不演示了。

	context.HTML(http.StatusOK, "login-after.tmpl", gin.H{
		"username": username,
		"password": password,
	})

}
