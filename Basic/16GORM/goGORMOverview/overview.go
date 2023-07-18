package goGORMOverview

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

func Overview() {

	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:lcz19930316@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	// 创建表(表名称就是结构体名称加个s后缀)
	db.AutoMigrate(&Product{})

	// 插入数据
	product1 := Product{
		Code:  "1001",
		Price: 100,
	}
	db.Create(&product1)

	// 查询数据
	// 首先声明一个 product2 用来接收查到的数据
	var product2 Product
	// 根据主键ID查询(First和Find两个方法都可以)
	// db.Find(&product2, 1)
	// 根据条件查询(查询code为 1001 的数据)
	db.Find(&product2, "code = ?", "1001")
	fmt.Printf("product2: %v\n", product2)

	// 更新数据
	// 首先声明一个 product3 用来接收查到的数据
	var product3 Product
	// 首先将要更新的数据查出来
	db.First(&product3, 1)
	// 更新单个字段
	db.Model(&product3).Update("price", 999)
	// 更新多个字段
	db.First(&product3, 2)
	fmt.Printf("product3: %v\n", product3)
	db.Model(&product3).Updates(Product{Price: 888, Code: "111"})
	// 或
	db.First(&product3, 5)
	db.Model(&product3).Updates(map[string]interface{}{"Price": 666, "Code": "222"})

	// 删除数据(这个删除并不是物理删除，实际上是给deleted_at字段的值添加了一个标记，进行一个软删除)
	// 首先声明一个 product3 用来接收查到的数据
	var product4 Product
	db.First(&product4, "code = ?", "1002")
	db.Delete(&product4, "code = ?", "1002")

}
