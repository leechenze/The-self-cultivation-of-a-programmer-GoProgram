package goInitialDataBase

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

func GoInitialDataBase() (db *sql.DB, err error) {
	dsn := "root:lcz19930316@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		return nil, err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// 打印数据库
	fmt.Printf("&db: %v\n", db)
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
	return db, nil
}
