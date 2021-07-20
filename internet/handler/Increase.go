package handler

import (
	"awesomeProject/Testfourth/internet/mondel"
	"awesomeProject/Testfourth/response"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go.mongodb.org/mongo-driver/bson"
)

//用户增加奖励
func Increase(id string,temp []byte) []byte {
	tmp := make(map[string]interface{})
	err := json.Unmarshal(temp, &tmp)
	fmt.Println("tmp:",tmp)
	if err!=nil{
		fmt.Println(err)
	}
	tmpPack := mondel.Pack{}
	//取出data数据
	bytes, err := json.Marshal(tmp["Data"])
	json.Unmarshal(bytes,&tmpPack)
	fmt.Println("tmppack:",tmpPack)
	file := bson.D{{"id",id}}
	result := mondel.User{}
	database := Client.Database("user")
	collection := database.Collection("user")
	collection.FindOne(context.TODO(), file).Decode(&result)

	//让礼品码内容和用户余额相加
	newGold := tmpPack.Gold + result.Gold
	newDiamond := tmpPack.Diamond + result.Diamond
	updateResult, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.M{"$set": bson.M{"gold": newGold, "diamond": newDiamond}},
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
	marshal, err := proto.Marshal(ur)
	if err!=nil{
		fmt.Println(err)
	}
	return marshal
}
