package dataBaseActionFunc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func UpdateMany(client *mongo.Client) {
	collection := client.Database("go_db").Collection("student")

	bsonFilter := bson.D{{"name", "tom1"}}
	bsonCondition := bson.D{{
		"$set",
		bson.D{
			{"name", "nixon"},
			{"age", 22},
		},
	}}
	res, err := collection.UpdateMany(context.TODO(), bsonFilter, bsonCondition)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Printf("res.ModifiedCount: %v\n", res.ModifiedCount)

}
