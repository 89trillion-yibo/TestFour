package main

import (
	"awesomeProject/Testfourth/internet/handler"
	"awesomeProject/Testfourth/response"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestIncrease(t *testing.T)  {
	var exp string
	exp = "15"
	user, isBool, errmessage := handler.JudgmentUser(exp)
	fmt.Println(user,isBool,errmessage)
}

func TestInfunc(t *testing.T) {
	increase, errmessage := handler.Increase("1", "V5UF7NR9")
	newStu := response.GeneralReward{}
	proto.Unmarshal(increase,&newStu)
	fmt.Println(newStu)
	fmt.Println(errmessage)
}
