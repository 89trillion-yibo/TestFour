package main

import (
	"awesomeProject/fourth/service"
	"fmt"
	"testing"
)

func TestIncrease(t *testing.T)  {
	var exp string
	exp = "15"
	user := service.JudgmentUser(exp)
	fmt.Println(user)
}
