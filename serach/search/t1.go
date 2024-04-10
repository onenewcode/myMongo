package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

// Restaurant 结构作为 restaurants 集合中文档的模型
type Restaurant struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string
	RestaurantId string `bson:"restaurant_id"`
	Cuisine      string
	Address      interface{}
	Borough      string
	Grades       []interface{}
}

func main() {
	// 设置uri环境变量用于设置MONGODB地址,这里我们使用，官网提供的云服务
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
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_restaurants").Collection("restaurants")
	//创建查询筛选器以匹配匹配“name”是"Bagels N Buns"的文档
	filter := bson.D{{"name", "Bagels N Buns"}}
	// 检索第一个匹配的文档
	var result Restaurant
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	// 	如果没有匹配的文档或有文档，则打印消息，
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return
		}
		panic(err)
	}
	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
}
