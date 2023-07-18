package dataBaseActionFunc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Query(client *mongo.Client) {
	collection := client.Database("go_db").Collection("student")
	// bson.D{} 表示查找的过滤条件为查询所有，返回cursor游标
	// cursor, err := collection.Find(context.TODO(), bson.D{})
	// 声明查询条件的name为tom，并且age为20。
	cursor, err := collection.Find(context.TODO(), bson.D{{"name", "tom"}, {"age", 20}})
	// 最后关闭游标
	defer cursor.Close(context.TODO())
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	// 遍历游标
	for cursor.Next(context.TODO()) {
		var result bson.D
		cursor.Decode(&result)

		fmt.Printf("result: %v\n", result)
	}

}
