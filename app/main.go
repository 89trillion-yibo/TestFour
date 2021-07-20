package main

import (
	"awesomeProject/Testfourth/app/http"
	"awesomeProject/Testfourth/internet/handler"
	"awesomeProject/Testfourth/internet/mondel"
	"fmt"
)

func main() {

	//初始化mongodb连接
	handler.Client = mondel.MangodbConnection()
	//启动服务
	err := http.InitRun()
	if err!=nil{
		fmt.Println(err)
	}

	/*engine:= gin.Default()
    engine.POST("/getUser", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		user := handler.JudgmentUser(uuid)
		if len(user) != 0 {
			c.JSON(200,user)
		}else {
			c.JSON(200,"信息错误")
		}
	})

	engine.POST("/increase", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		gifcode := c.PostForm("gifcode")
		user := handler.JudgmentUser(uuid)
		if len(user) != 0 {
			bytes := internet.HttpClient(uuid, gifcode)
			increase := handler.Increase(uuid, bytes)
			c.JSON(200,increase)
		}else {
			c.JSON(200,"信息错误")
		}
	})
	engine.Run(":8081")*/
}
