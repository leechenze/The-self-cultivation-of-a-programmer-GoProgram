package controller

import (
	"GoWeb/GinToDoList/dao"
	"GoWeb/GinToDoList/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateToDo(ctx *gin.Context) {
	var todo model.Todo
	db, _ := dao.InitDataBase()

	// 获取路径参数 /:id
	id := ctx.Param("id")
	// 查询对应ID的数据
	err := db.Where("id = ?", id).First(&todo).Error
	// 可以根据影响的行数(RowsAffected)进行判断，也可以通过返回错误进行判断
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":      500,
			"message":   "操作失败",
			"errorInfo": err.Error(),
		})
		return
	} else {
		// 绑定参数（将前段传过来的参数绑定到 todo对象中）
		ctx.ShouldBind(&todo)
		// 更新数据,返回错误判断的简写方式
		if err := db.Save(&todo).Error; err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code":      500,
				"message":   "操作失败",
				"errorInfo": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, todo)
		}
	}
}
