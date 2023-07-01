package dataBaseActionFunc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type student struct {
	Name string
	Age  int
}

func Insert(client *mongo.Client) {
	s1 := student{
		Name: "tom",
		Age:  20,
	}

	collection := client.Database("go_db").Collection("student")
	res, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("res.InsertedID: %v\n", res.InsertedID)
	}

}

func InsertMany(client *mongo.Client) {
	collection := client.Database("go_db").Collection("student")
	s1 := student{
		Name: "hillary",
		Age:  20,
	}
	s2 := student{
		Name: "clinton",
		Age:  20,
	}

	// []interface{} 表示任意类型
	res, err := collection.InsertMany(context.TODO(), []interface{}{s1, s2})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("res.InsertedIDs: %v\n", res.InsertedIDs)
	}

}
