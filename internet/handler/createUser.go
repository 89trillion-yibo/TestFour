package handler

import (
	"awesomeProject/Testfourth/internet/model"
	"awesomeProject/Testfourth/internet/service"
)

//判断用户是否存在，不存在则注册创建
func JudgmentUser(uuid string) (model.User,bool,string) {
	var user model.User
	//判断用户是否存在
	exist, errString := service.IsHaveUser(uuid)
	if errString != "" {
		return user,false,errString
	}
	//存在则登陆
	if exist {
		user,errString = model.Findone(uuid)
		if errString != "" {
			return user,false,errString
		}
		return user,true,""
	}else {
		//不存在则注册
		user.Id = uuid
		user.Gold = 0
		user.Diamond = 0
		insertOK, err := model.Insertone(user)
		if insertOK {
			return user,true,err
		}
		return user,false,err
	}
}




