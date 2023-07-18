package GinToDoList

import (
	"GoWeb/GinToDoList/router"
)

func Main() {
	// 连接数据库
	// db, err := dao.InitDataBase()
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	// 创建数据表(只需要创建一次即可)
	// 	db.AutoMigrate(&model.Todo{})
	// }

	// 路由层拆分
	engine := router.SetupRouter()
	engine.Run(":6900")

}
