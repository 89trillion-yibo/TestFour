package router

import (
	"awesomeProject/Testfourth/internet/ctrl"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) *gin.Engine {
	//用户登陆注册
	r.POST("/getUser", ctrl.CreateUser)
	//用户添加奖励
	r.POST("/increase",ctrl.IncreaseGife)
	//管理员添加礼品码
	r.POST("/creategif", ctrl.CreateCode)
	//管理员查找礼品码
	r.POST("/getGifcode",ctrl.GetGifcode)
	return r
}