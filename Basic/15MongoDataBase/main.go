package main

import (
	"context"
	"log"
	"mongoDataBase/dataBaseActionFunc"
)

func main() {
	println("========================Mongo DataBase========================")
	// 初始化Mongo连接
	println()
	mongoClient, err := dataBaseActionFunc.InitConnect()
	// 关闭客户端连接
	defer mongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatalf("数据库连接失败: %v\n", err)
	}

	// bson
	println()
	// dataBaseActionFunc.BsonFn()

	// insert
	println()
	// dataBaseActionFunc.Insert(mongoClient)
	// dataBaseActionFunc.InsertMany(mongoClient)

	// query
	println()
	// dataBaseActionFunc.Query(mongoClient)

	// update
	println()
	// dataBaseActionFunc.UpdateMany(mongoClient)

	// delete
	println()
	dataBaseActionFunc.Delete(mongoClient)
}
