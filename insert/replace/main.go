// Replaces the first document that matches a filter by using the Go driver
package main

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// start-restaurant-struct
type Restaurant struct {
	Name         string
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}

// end-restaurant-struct

func main() {
	err := os.Setenv("MONGODB_URI", "mongodb+srv://minioadmin:minioadmin@cluster0.f22y2qo.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
	if err != nil {
		return
	}
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_restaurants").Collection("restaurants")
	// 生成过滤器
	filter := bson.D{{"name", "Madame Vo"}}

	replacement := Restaurant{Name: "Monsieur Vo", Cuisine: "Asian Fusion"}

	// 将与筛选器匹配的第一个文档替换为新文档
	result, err := coll.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		panic(err)
	}

	if result.MatchedCount != 0 {
		fmt.Println("Number of documents replaced: %d\n", result.ModifiedCount)
	}

}
