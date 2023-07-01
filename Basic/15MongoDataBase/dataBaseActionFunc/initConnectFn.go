package dataBaseActionFunc

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitConnect() (client *mongo.Client, err error) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err1 := conn.Ping(context.TODO(), nil)
	if err1 != nil {
		return nil, err1
	}

	log.Println("连接成功")
	return conn, nil
}
