package service

import "awesomeProject/Testfourth/internet/model"

//判断用户是否存在
func IsHaveUser(uuid string) (bool,string) {
	exist, errMessage := model.Exist(uuid)
	return exist,errMessage
}

