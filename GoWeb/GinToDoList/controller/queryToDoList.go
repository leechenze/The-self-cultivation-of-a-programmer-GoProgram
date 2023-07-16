package controller

import (
	"GoWeb/GinToDoList/dao"
	"GoWeb/GinToDoList/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func QueryToDoList(ctx *gin.Context) {
	var todoList []model.Todo
	db, _ := dao.InitDataBase()

	// 查询所有的待办事项
	tx := db.Find(&todoList)
	// 转换前段传过来的参数数据进行打印
	marshal, _ := json.Marshal(todoList)
	fmt.Printf("GET return result:%v\n", string(marshal))
	// 根据影响的行数进行判断
	if tx.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}
