package GinParams

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func BindGetParams(context *gin.Context) {
	// username := context.Query("username")
	// password := context.Query("password")
	// userinfo := UserInfo{
	// 	username: username,
	// 	password: password,
	// }

	var userInfo UserInfo

	// 绑定参数
	// 切记：如果在一个函数或方法中修改传进去的值，就必须传递这个值的指针。
	// 然后这个值的结构体必须是大写也就是值是可访问的。
	err := context.ShouldBind(&userInfo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("get userInfo: %v\n", userInfo)
		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func BindPostParams(context *gin.Context) {

	// 这里用postMan测试一下

	var userInfo UserInfo

	// 绑定参数
	err := context.ShouldBind(&userInfo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("psot userInfo: %v\n", userInfo)
		context.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}
