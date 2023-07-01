package dataBaseActionFunc

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func BsonFn() {
	doc := bson.D{{"name", "tom"}}
	fmt.Printf("doc: %v\n", doc)
}
