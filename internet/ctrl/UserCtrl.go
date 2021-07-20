package ctrl

import (
	"awesomeProject/Testfourth/internet"
	"awesomeProject/Testfourth/internet/gifeerror"
	"awesomeProject/Testfourth/internet/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查找用户或创建用户
func CreateUser(c *gin.Context)  {
	uuid, b := c.GetPostForm("uuid")
	if !b {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}
	user:= handler.JudgmentUser(uuid)
	if len(user) != 0  {
		//mongodb已有用户
		c.JSON(http.StatusOK,gifeerror.HasUser.AddData(user))
	}else {
		//mongodb没有该用户，创建新用户
		c.JSON(http.StatusOK,gifeerror.NoHasUser.AddData(user))
	}
}

//查找用户并增加奖励
func IncreaseGife(c *gin.Context)  {
	uuid,a := c.GetPostForm("uuid")
	gifcode,b := c.GetPostForm("gifcode")
	if !a || !b {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}
	user := handler.JudgmentUser(uuid)
	if len(user) != 0  {
		bytes := internet.HttpClient(uuid, gifcode)
		increase := handler.Increase(uuid, bytes)
		c.JSON(http.StatusOK,gifeerror.HasUser.AddData(increase))
	}else if len(user) == 0 {
		bytes := internet.HttpClient(uuid, gifcode)
		increase := handler.Increase(uuid, bytes)
		c.JSON(http.StatusOK,gifeerror.NoHasUser.AddData(increase))
	} else {
		c.JSON(http.StatusBadGateway,gifeerror.Error)
	}
}