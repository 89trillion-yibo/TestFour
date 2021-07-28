package ctrl

import (
	"awesomeProject/Testfourth/internet/gifeerror"
	"awesomeProject/Testfourth/internet/handler"
	"awesomeProject/Testfourth/internet/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

//查找用户或创建用户
func CreateUser(c *gin.Context)  {
	uuid, b := c.GetPostForm("uuid")
	if !b || uuid == "" {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		user, isBool, errMessage := handler.JudgmentUser(uuid)
		if isBool  {
			//mongodb已有用户或已经注册
			c.JSON(http.StatusOK,gifeerror.HasUser.AddData(user))
		}else {
			//发生错误
			c.JSON(http.StatusBadGateway,gifeerror.SQLError.AddMessage(errMessage))
		}
	}
}

//查找用户并增加奖励
func IncreaseGife(c *gin.Context)  {
	uuid,a := c.GetPostForm("uuid")
	gifcode,b := c.GetPostForm("gifcode")
	if !a || !b {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}
	//先判断是否有该用户
	isHave,errMessage := service.IsHaveUser(uuid)
	if isHave  {
		//如果有该用户的话再判断是否有领取资格
		judgment, err := service.Judgment(uuid, gifcode)
		if judgment {
			//如果有资格则增加奖励返回数据
			increase,errmessage := handler.Increase(uuid, gifcode)
			//如果有错误消息，说明修改未成功，回滚数据
			if errmessage!=""{
				c.JSON(http.StatusBadGateway,gifeerror.SQLError.AddMessage(errmessage))
			}else {
				//没有报错消息则输出数据
				c.JSON(http.StatusOK,gifeerror.OK.AddData(increase))
			}
		}else {
			//没有领取资格返回原因
			c.JSON(http.StatusBadGateway,gifeerror.SQLError.AddMessage(err))
		}
	}else if errMessage!= "" {
		//如果判断是否有该用户时出现错误，返回错误信息
		c.JSON(http.StatusBadGateway,gifeerror.SQLError.AddMessage(errMessage))
	}else {
		//如果没有该用户则先创建用户
		c.JSON(http.StatusOK,gifeerror.CreateFirst)
	}
}