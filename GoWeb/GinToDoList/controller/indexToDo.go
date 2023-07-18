package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(ctx *gin.Context) {
	// index.html 就是加载的模版文件的下的 index.html
	ctx.HTML(http.StatusOK, "index.html", nil)
}
