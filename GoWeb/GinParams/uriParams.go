package GinParams

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UriParams(context *gin.Context) {

	// 以 http://localhost:9090/leechenze/100 这个路径为例

	// 获取路径参数(对应：/:name/:age )
	name := context.Param("name")
	age := context.Param("age")
	context.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}
