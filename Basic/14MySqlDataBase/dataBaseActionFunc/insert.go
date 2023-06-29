package dataBaseActionFunc

import (
	"database/sql"
	"fmt"
)

func Insert(db *sql.DB) {
	sqlStatement := "insert into user_tbl (username, password) values (?,?)"

	res, err := db.Exec(sqlStatement, "hillary", "123")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		id, _ := res.LastInsertId()
		fmt.Printf("res.LastInsertId: %v\n", id)
	}
}
