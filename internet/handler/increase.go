package handler

import (
	"awesomeProject/Testfourth/internet/model"
	"awesomeProject/Testfourth/response"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
)

//用户增加奖励
func Increase(id string,gifecode string) ([]byte,string) {
	//获取奖励内容
	get := model.HashGet(gifecode, "Pack")
	bytes := []byte(get)
	tmp := make(map[string]interface{})
	err := json.Unmarshal(bytes, &tmp)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("tmp>>>>>>>>>>",tmp)

	//获取用户信息
	user, errmessage := model.Findone(id)
	if errmessage != "" {
		return nil,errmessage
	}
	//添加奖励
	userPastGold := float64(user.Gold)
	userPastDiamond := float64(user.Diamond)
	newGold := int(userPastGold + tmp["Gold"].(float64))
	newDiamond := int(userPastDiamond + tmp["Diamond"].(float64))
	user.Gold = newGold
	user.Diamond = newDiamond
	//修改用户数据
	updataOK, errmessage := model.Updataone(user)
	if !updataOK {
		//defer实现数据回滚
		defer func() {
			decr := model.Decr(gifecode + ":Bytime")
			fmt.Println("decr:",decr)
			hdel := model.Hdel(gifecode+":receive", id)
			fmt.Println("hdel:",hdel)
		}()
		return nil,errmessage
	}
	ur := response.GeneralReward{
		Code: 200,
		Msg: "信息返回正确,1001代表金币,1002代表钻石",
		Changes: map[uint32]uint64{
			1001:uint64(tmp["Gold"].(float64)),
			1002:uint64(tmp["Diamond"].(float64)),
		},
		Balance: map[uint32]uint64{
			1001:uint64(userPastGold),
			1002:uint64(userPastDiamond),
		},
		Counter: map[uint32]uint64{
			1001:uint64(user.Gold),
			1002:uint64(user.Diamond),
		},
	}
	fmt.Println("ur>>>>>>>>>>>>>:",ur)
	marshal, err := proto.Marshal(&ur)
	if err!=nil{
		fmt.Println(err)
	}
	return marshal,""
}


