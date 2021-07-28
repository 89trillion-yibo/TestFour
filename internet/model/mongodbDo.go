package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

//判断信息是否存在
func Exist(id string) (bool,string)  {
	var user User
	err := Conn.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&user)
	if user.Id != "" {
		return true,""
	}
	if err!=nil && err.Error() != "mongo: no documents in result"{
		return false,"查找数据库发生错误"
	}
	return false,""
}

//添加单条信息
func Insertone(user User) (bool,string) {
	_, err := Conn.InsertOne(context.TODO(), user)
	if err!=nil{
		fmt.Println(err)
		return false,"插入发生错误回滚"
	}
	return true,""
}

//查找信息
func Findone(id string) (User,string) {
	var user User
	err := Conn.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&user)
	if err!=nil && err.Error() != "mongo: no documents in result"{
		return user,"查找数据库发生错误"
	}
	return user,""
}

//修改信息
func Updataone(user User) (bool,string) {
	_, err := Conn.UpdateOne(context.TODO(), bson.D{{"id", user.Id}}, bson.D{{"$set", user}})
	if err!=nil {
		return false,"修改出错回滚数据"
	}
	return true,""
}