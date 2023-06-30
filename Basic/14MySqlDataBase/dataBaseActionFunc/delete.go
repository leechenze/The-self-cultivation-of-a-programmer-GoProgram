package dataBaseActionFunc

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB) {

	sqlStatement := "delete from user_tbl where id = ?"

	res, err := db.Exec(sqlStatement, 4)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	affectedLine, _ := res.RowsAffected()
	fmt.Printf("affectedLine: %v\n", affectedLine)

}
