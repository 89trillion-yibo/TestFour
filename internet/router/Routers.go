package router

import (
	"awesomeProject/Testfourth/internet/ctrl"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.Engine) *gin.Engine {
	r.POST("/getUser", ctrl.CreateUser)
	r.POST("/increase",ctrl.IncreaseGife)
	return r
}