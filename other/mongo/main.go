package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// 字典结构体
type Dictionary struct {
	Id               primitive.ObjectID `bson:"_id"`
	ItemCode         string             `bson:"item_code"`
	ItemValue        string             `bson:"item_value"`
	ItemName         string             `bson:"item_name"`
	LastModifiedTime string             `bson:"last_modified_time"`
	ItemStatus       int32              `bson:"item_status"`
	DisplayOrder     string             `bson:"display_order"`
	ItemType         string             `bson:"item_type"`
}

func main() {
	fmt.Println("mongodb")
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://10.5.7.122:27018")
	//  连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDb!")
	// 指定获取要操作的数据集
	getItemListByType(client, "company_property")

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func getItemListByType(client *mongo.Client, typeName string) {
	collection := client.Database("b2b_parameter").Collection("infr_dictionary")
	//var result Dictionary
	//filter := bson.D{{"item_type",typeName},{"item_code","1"}}
	//err := collection.FindOne(context.Background(), filter).Decode(&result)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 获取总数
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("countDocuments: ", count)
	// 获取多条记录(游标)
	cur, err := collection.Find(context.Background(), bson.M{"item_type": typeName})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("documents: %+v\n", cur)
	//var all []*Dictionary
	//err = cur.All(context.Background(), &all)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("collection.Find curl.All: ", all)
	for cur.Next(context.Background()) {
		var b Dictionary
		if err = cur.Decode(&b); err != nil {
			log.Fatal(err)
		}
		fmt.Println("cur.Next: ", b)
		fmt.Println("cur.Next.itemName", b.ItemName)
	}
	_ = cur.Close(context.Background())
	//fmt.Println("collection.FindOne", result)
	//fmt.Printf("Found a single document: %+v\n", result)
	//fmt.Printf("item name is %s\n", result.item_name)
}
