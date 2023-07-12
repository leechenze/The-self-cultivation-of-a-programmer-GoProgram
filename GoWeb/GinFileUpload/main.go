package GinFileUpload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func Main() {
	engine := gin.Default()
	engine.LoadHTMLFiles("./GinFileUpload/index.tmpl")
	engine.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.tmpl", nil)
	})

	engine.POST("/upload", func(context *gin.Context) {
		// 从请求中读取文件
		file, err := context.FormFile("file1")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		} else {
			dst := path.Join("./GinFileUpload", file.Filename)
			fmt.Printf("dst: %v\n", dst)
			// 将读取到的文件保存到本地（服务端本地）
			_ = context.SaveUploadedFile(file, dst)
			context.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})

	// 多文件上传同理，只不过读取文件的方法改为了：MultipartForm()，返回的是多个文件，循环它保存每一个即可。

	engine.Run(":9090")
}
