package service

import (
	"awesomeProject/Testfourth/connection"
	"awesomeProject/Testfourth/pojo"
	"awesomeProject/Testfourth/response"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

var client = connection.MangodbConnection()

//判断用户是否存在，不存在则注册创建
func JudgmentUser(uuid string) map[string]interface{} {
	file := bson.D{{"Id",uuid}}
	result := make(map[string]interface{})
	database := client.Database("user")
	collection := database.Collection("user")
	err := collection.FindOne(context.TODO(), file).Decode(&result)
	if err!=nil{
		fmt.Println(err)
		tmp := pojo.User{Id: uuid,Gold: 0,Diamond: 0}
		insertOne, err := collection.InsertOne(context.TODO(), tmp)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(insertOne.InsertedID)
	}
	return result
}

//用户增加奖励
func Increase(id string,temp []byte) response.GeneralReward {
	tmpPack := pojo.Pack{}
	err := json.Unmarshal(temp, &tmpPack)
	if err!=nil{
		fmt.Println(err)
	}
	file := bson.D{{"Id",id}}
	result := pojo.User{}
	database := client.Database("user")
	collection := database.Collection("user")
	collection.FindOne(context.TODO(), file).Decode(&result)

	newGold := tmpPack.Gold + result.Gold
	newDiamond := tmpPack.Diamond + result.Diamond
	updateResult, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"Id": id},
		bson.M{"$set": bson.M{"Gold": newGold, "Diamond": newDiamond}},
	)
	var ur = &response.GeneralReward{}
	if updateResult.MatchedCount == 1 {
		 ur = &response.GeneralReward{
			Code: 200,
			Msg: "信息返回正确,1001代表金币,1002代表钻石",
			Changes: map[uint32]uint64{
				1001:uint64(tmpPack.Gold),
				1002:uint64(tmpPack.Diamond),
			},
			Balance: map[uint32]uint64{
				1001:uint64(result.Gold),
				1002:uint64(result.Diamond),
			},
			Counter: map[uint32]uint64{
				1001:uint64(newGold),
				1002:uint64(newDiamond),
			},
			Ext: "扩展字段",
		}
	}
	fmt.Println(*ur)
	return *ur
}
