package goGORMActionFunc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDataBase() (dataBase *gorm.DB, err error) {
	dsn := "root:lcz19930316@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:               dsn,
				DefaultStringSize: 256, // string字段类型的默认长度
			},
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 打印Sql日志
		},
	)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}
	return db, nil
}
