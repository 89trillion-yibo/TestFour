package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Conn *mongo.Collection
var Session  mongo.Session

func MangodbConnection()  {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//连接数据库
	database := client.Database("user")
	Conn = database.Collection("user")
	//开启事务
	session, err := client.StartSession()
	if err != nil {
		fmt.Println(err)
	}
	Session = session
}
