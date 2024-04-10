package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
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

	// 构建插入对象
	newRestaurants := []interface{}{
		Restaurant{Name: "Rule of Thirds", Cuisine: "Japanese"},
		Restaurant{Name: "Madame Vo", Cuisine: "Vietnamese"},
	}

	// 插入多行数据到集合
	result, err := coll.InsertMany(context.TODO(), newRestaurants)
	if err != nil {
		panic(err)
	}

	// 输出新生成的id
	fmt.Printf("%d documents inserted with IDs:\n", len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		fmt.Printf("\t%s\n", id)
	}

}
