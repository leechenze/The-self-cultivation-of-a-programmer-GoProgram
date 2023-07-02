package main

import (
	"fmt"
	"gorm/goGORMActionFunc"
)

func main() {
	println("========================GORM========================")
	// GORM概述与初探：
	// println()
	// goGORMOverview.Overview()

	// 连接数据库：
	println()
	db, err := goGORMActionFunc.InitDataBase()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 创建记录：
	// println()
	// goGORMActionFunc.Create(db)

	// 查询记录：
	// println()
	// goGORMActionFunc.Query(db)

	// 更新记录：
	// println()
	// goGORMActionFunc.Update(db)

	// 删除记录：
	// println()
	// goGORMActionFunc.Delete(db)

	// 原生Sql和Sql构建器：
	// println()
	goGORMActionFunc.SqlBuilder(db)

}
