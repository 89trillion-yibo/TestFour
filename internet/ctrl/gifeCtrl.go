package ctrl

import (
	"awesomeProject/Testfourth/internet/gifeerror"
	"awesomeProject/Testfourth/internet/model"
	"awesomeProject/Testfourth/internet/service"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//创建礼品码
func CreateCode(c *gin.Context)  {

	//接收参数
	gifType, a := c.GetPostForm("gifType")
	des , b := c.GetPostForm("des")
	allowTime, d := c.GetPostForm("allowTime")
	valTime, e := c.GetPostForm("valTime")
	createName, f := c.GetPostForm("createName")
	gold, g := c.GetPostForm("gold")
	diamond, h := c.GetPostForm("diamond")
	//判断是否有参数
	if !a || !b || !d || !e || !f || !g || !h {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//将string转成int
		giftype, _ := strconv.Atoi(gifType)
		alltime, _ := strconv.Atoi(allowTime)
		jinbi, _ := strconv.Atoi(gold)
		zuanshi, _ := strconv.Atoi(diamond)
		//初始化礼品内容
		s := model.Pack{Gold: jinbi,Diamond: zuanshi}
		marshal, _ := json.Marshal(s)
		pack := string(marshal)
		//创建礼品码
		careatGif, mapdate := service.CareatGif(giftype, des, alltime, valTime, pack, createName)
		fmt.Println("code:",mapdate)

		c.JSON(http.StatusOK,gifeerror.OK.AddData(careatGif))
	}
}

//查询礼品码
func GetGifcode(c *gin.Context) {
	//接收参数
	gifcode, a := c.GetPostForm("gifcode")
	if !a || gifcode == "" {
		c.JSON(http.StatusBadRequest,gifeerror.Parameters)
	}else {
		//获取礼品码信息
		data, err := model.HashGetAll(gifcode)
		//获取领取列表信息
		receive, err := model.HashGetAll(gifcode + ":receive")
		//获取已领取次数
		bytime, err := model.StringGet(gifcode + ":Bytime")
		gifeAndReceive := make(map[string]interface{})
		gifeAndReceive["gifeInfo"] = data
		gifeAndReceive["receiveInfo"] = receive
		gifeAndReceive["bytimeInfo"] = bytime
		if err!="" {
			c.JSON(http.StatusBadGateway,gifeerror.SQLError.AddMessage(err))
		}
		c.JSON(http.StatusOK,gifeerror.OK.AddData(gifeAndReceive))
	}

}
