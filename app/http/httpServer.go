package http

import (
	"awesomeProject/Testfourth/internet/router"
	"github.com/gin-gonic/gin"
)

func InitRun() error {
	r := gin.Default()
	//启动路由
	router.Routers(r)
	err := r.Run(":8081")
	if err!=nil{
		return err
	}
	return nil
}