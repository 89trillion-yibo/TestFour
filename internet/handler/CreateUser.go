package handler

import (
	"awesomeProject/Testfourth/internet/mondel"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

//判断用户是否存在，不存在则注册创建
func JudgmentUser(uuid string) (map[string]interface{}) {
	file := bson.D{{"id",uuid}}
	result := make(map[string]interface{})
	database := Client.Database("user")
	collection := database.Collection("user")
	err := collection.FindOne(context.TODO(), file).Decode(&result)
	if err!=nil{
		fmt.Println(err)
		tmp := mondel.User{Id: uuid,Gold: 0,Diamond: 0}
		insertOne, err := collection.InsertOne(context.TODO(), tmp)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(insertOne.InsertedID)
	}
	return result
}


