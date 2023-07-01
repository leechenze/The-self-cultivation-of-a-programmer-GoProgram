package dataBaseActionFunc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Delete(client *mongo.Client) {

	collection := client.Database("go_db").Collection("student")

	bsonDelFilter := bson.D{{"name", "tom1"}}

	res, err := collection.DeleteMany(context.TODO(), bsonDelFilter)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	fmt.Printf("res.DeletedCount: %v\n", res.DeletedCount)

}
