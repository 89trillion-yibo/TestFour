package mondel

//用户结构
type User struct {
	Id string    `json:"id"`      //用户id
	_id string   `json:"_id"`   //mongodb自动生成的id
	Gold int     `json:"gold"`   //金币
	Diamond int  `json:"diamond"`    //钻石
}

//奖励结构
type Pack struct {
	Gold int      `json:"Gold"`    //金币
	Diamond int   `json:"Diamond"` //钻石
}
