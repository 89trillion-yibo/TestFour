package pojo

type User struct {
	Id string
	_id string
	Gold int
	Diamond int
}

type Pack struct {
	Gold int      `json:"Gold"`
	Diamond int   `json:"Diamond"`
}
