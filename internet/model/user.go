package model

//用户结构
type User struct {
	Id string    `json:"id"`      //用户id
	_id string   `json:"_id"`   //mongodb自动生成的id
	Gold int     `json:"gold"`   //金币
	Diamond int  `json:"diamond"`    //钻石
}
