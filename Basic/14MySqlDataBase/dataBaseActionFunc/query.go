package dataBaseActionFunc

import (
	"database/sql"
	"fmt"
)

type User struct {
	id       int
	username string
	password string
}

var user User

// 单行查询
func Query(db *sql.DB) {
	sqlStatement := "select * from user_tbl where id = ?"

	// 参数 1 对应 sql语句中的 ?，Scan 方法表示将查到的值赋给那个对象。
	err := db.QueryRow(sqlStatement, 1).Scan(&user.id, &user.username, &user.password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("user: %v\n", user)
	}
}

// 多行查询
func MultilineQuery(db *sql.DB) {
	sqlStatement := "select * from user_tbl"
	rows, err := db.Query(sqlStatement)
	// 关闭rows释放持有的数据库连接
	defer rows.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		for rows.Next() {
			err := rows.Scan(&user.id, &user.username, &user.password)
			if err != nil {
				fmt.Printf("err: %v\n", err)
			} else {
				fmt.Printf("user: %v\n", user)
			}
		}
	}
}
