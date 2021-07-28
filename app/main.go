package main

import (
	"awesomeProject/Testfourth/app/http"
	"awesomeProject/Testfourth/internet/model"
	"fmt"
)

func main() {
	//初始化mongodb连接
	model.MangodbConnection()
	//初始化redis连接
	model.RedisCli("127.0.0.1:6379")
	//启动服务
	err := http.InitRun()
	if err!=nil{
		fmt.Println(err)
	}
}
