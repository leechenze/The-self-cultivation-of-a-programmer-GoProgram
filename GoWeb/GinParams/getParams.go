package GinParams

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetParams(context *gin.Context) {

	// 本章示例以 http://localhost:9090/?name=林肯 为准

	// queryVal := context.Query("name")
	// queryVal := context.DefaultQuery("name", "name is nil")
	queryVal, ok := context.GetQuery("name")
	if !ok {
		queryVal = "name is nil"
	}
	context.JSON(http.StatusOK, gin.H{
		"name": queryVal,
	})

}
