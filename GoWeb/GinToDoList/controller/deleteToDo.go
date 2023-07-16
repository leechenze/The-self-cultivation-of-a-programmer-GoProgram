package controller

import (
	"GoWeb/GinToDoList/dao"
	"GoWeb/GinToDoList/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteToDo(ctx *gin.Context) {
	db, _ := dao.InitDataBase()

	id := ctx.Param("id")
	tx := db.Where("id = ?", id).Delete(model.Todo{})
	// 根据影响的行数进行判断
	if tx.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "删除成功",
		})
	}
}
