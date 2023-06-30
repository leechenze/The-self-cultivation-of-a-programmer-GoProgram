package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mysqldatabase/goInitialDataBase"
)

func main() {
	println("========================MySql DataBase========================")
	// 初始化连接数据库
	db, err := goInitialDataBase.GoInitialDataBase()
	// 自增ID重置：设置为1或0，也可达到重置自增id的目的，如果无用则取当前数据最后行+1或者最后id+1
	db.Exec("ALTER TABLE user_tbl AUTO_INCREMENT = 0")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		println("数据库连接成功")

		// 插入数据
		println()
		// defer dataBaseActionFunc.Insert(db)

		// 查询数据
		println()
		// defer dataBaseActionFunc.Query(db)
		// defer dataBaseActionFunc.MultilineQuery(db)

		// 更新数据
		println()
		// defer dataBaseActionFunc.Update(db)

		// 删除数据
		println()
		// dataBaseActionFunc.Delete(db)

	}

}
