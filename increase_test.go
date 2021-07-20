package main

import (
	service2 "awesomeProject/Testfourth/internet/handler"
	"awesomeProject/Testfourth/response"
	"encoding/json"
	"fmt"
	"testing"
)

func TestIncrease(t *testing.T)  {
	var exp string
	exp = "15"
	user := service2.JudgmentUser(exp)
	fmt.Println(user)
}

func TestInfunc(t *testing.T) {
	tmp := make(map[string]int)
	tmp["Gold"] = 1000
	tmp["Diamond"] = 100
	marshal, _ := json.Marshal(tmp)
	increase := service2.Increase("17", marshal)
	newStu := response.GeneralReward{}
	json.Unmarshal(increase,&newStu)
	fmt.Println(newStu)
}
