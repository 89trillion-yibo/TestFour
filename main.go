package main

import (
	"awesomeProject/Testfourth/internet"
	"awesomeProject/Testfourth/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func main() {
	engine:= gin.Default()
    engine.POST("/getUser", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		user := service.JudgmentUser(uuid)
		if len(user) != 0 {
			c.JSON(200,user)
		}else {
			c.JSON(200,"信息错误")
		}
	})

	engine.POST("/increase", func(c *gin.Context) {
		uuid := c.PostForm("uuid")
		gifcode := c.PostForm("gifcode")
		user := service.JudgmentUser(uuid)
		if len(user) != 0 {
			bytes := internet.HttpClient(uuid, gifcode)
			increase := service.Increase(uuid, bytes)
			i, _ := json.Marshal(increase)
			c.JSON(200,i)
		}else {
			c.JSON(200,"信息错误")
		}
	})
	engine.Run(":8081")
}
