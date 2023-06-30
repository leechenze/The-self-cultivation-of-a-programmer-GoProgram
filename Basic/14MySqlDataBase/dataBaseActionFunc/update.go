package dataBaseActionFunc

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB) {
	sqlStatement := "update user_tbl set username=?, password=? where id = ?"
	res, err := db.Exec(sqlStatement, "hillary", "789", 3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		// 返回影响的行数
		affectedLine, _ := res.RowsAffected()
		fmt.Printf("affectedLine: %v\n", affectedLine)
	}
}
