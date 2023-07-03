package main

import (
	"fmt"
	"gorm/goGORMActionFunc"
	"gorm/goGORMAssociation"
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
	// goGORMActionFunc.SqlBuilder(db)

	// 关联关系：
	// 多对一的关系：
	// println()
	// goGORMAssociation.BelongsTo(db)
	// 一对一的关系：
	// println()
	// goGORMAssociation.HasOne(db)
	// 一对多的关系：
	// println()
	// goGORMAssociation.HasMany(db)
	// 多对多的关系：
	// println()
	// goGORMAssociation.ManyToMany(db)
	// 实体关联：
	// println()
	goGORMAssociation.EntityAssociation(db)

}
