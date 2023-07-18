package GinJson

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {

	engine := gin.Default()
	engine.GET("/json", func(context *gin.Context) {
		data := gin.H{
			"name":    "lincoln",
			"message": "hello world!",
			"age":     88,
		}
		context.JSON(http.StatusOK, data)
	})

	engine.Run(":9090")

}
