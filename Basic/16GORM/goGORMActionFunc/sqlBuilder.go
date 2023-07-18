package goGORMActionFunc

import (
	"gorm.io/gorm"
)

type Result struct {
	ID   int
	Name string
	Age  int
}

func SqlBuilder(db *gorm.DB) {
	// var user User

	// Raw（一般用于查询，除查询外的其他操作都用 Exec）
	// db.Raw("select id, name, age from users where name = ?", "Elizabeth").Scan(&user)
	// fmt.Printf("user: %v\n", user)

	// Exec（直接更新）
	// db.Exec("update users set age = ? where name = ?", 60, "Hillary")

	// Named（命名参数）
	// db.Where("age = @customVariable", sql.Named("customVariable", 88)).Find(&user)
	// fmt.Printf("user: %v\n", user)

	// DryRun模式（用于生成Sql语句，并不执行）
	// statement := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	// sql := statement.SQL.String()
	// fmt.Printf("sql: %v\n", sql)

	// ToSQL（功能同DryRun一样）
	// sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Model(&User{}).Where("id = ?", 88).Limit(10).Order("age desc").Find(&[]User{})
	// })
	// fmt.Printf("sql: %v\n", sql)

	// Row（结果封装Row）
	// var name string
	// var age int
	// row := db.Table("users").Where("name = ?", "Clinton").Select("name", "age").Row()
	// row.Scan(&name, &age)
	// fmt.Printf("name: %v\n", name)
	// fmt.Printf("age: %v\n", age)
	// Rows
	// rows, err := db.Table("users").Where("age > ?", 0).Select("name", "age").Rows()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	for rows.Next() {
	// 		rows.Scan(&name, &age)
	// 		fmt.Printf("name: %v\n", name)
	// 		fmt.Printf("age: %v\n", age)
	// 	}
	// }

	// 扫描Rows到结构体中
	// rows, err := db.Model(&User{}).Where("age > ?", 0).Select("name, age").Rows()
	// if err != nil {
	// 	fmt.Printf("err: %v\n", err)
	// } else {
	// 	db.ScanRows(rows, &user)
	// }
	// fmt.Printf("user: %v\n", user)

}
