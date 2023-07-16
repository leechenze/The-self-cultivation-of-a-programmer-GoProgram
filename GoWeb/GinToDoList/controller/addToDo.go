package controller

import (
	"GoWeb/GinToDoList/dao"
	"GoWeb/GinToDoList/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddToDo(ctx *gin.Context) {
	var todo model.Todo
	db, _ := dao.InitDataBase()

	// 绑定参数（将前段传过来的参数绑定到 todo对象中）
	ctx.ShouldBind(&todo)
	// 转换前段传过来的参数数据进行打印
	marshal, _ := json.Marshal(todo)
	fmt.Printf("POST return result:%v\n", string(marshal))
	// 给数据库插入该数据
	tx := db.Create(&todo)
	// 根据影响的行数进行判断
	if tx.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "操作失败",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "添加成功",
		})
	}
}
