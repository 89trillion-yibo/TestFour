package gifeerror

var (
	OK = response(200,"ok")
	HasUser = response(200,"已有用户或新用户注册")
	NoHasUser = response(200,"新创建用户")
	Error = response(500,"error")

	Parameters = response(101,"获取参数错误")
	Exchange = response(102,"兑换错误")
	CreateFirst = response(103,"请先创建用户")
	SQLError = customize(104)

)

//异常结构
type GifeErr struct {
	Data    interface{}      //返回数据
	Code    int              //错误码
	Message string           //错误信息
}

//不返回数据
func response(code int , message string) *GifeErr{
	return &GifeErr{
		Code: code,
		Message: message,
		Data: nil,
	}
}

//自定义message
func customize(code int) *GifeErr {
	return &GifeErr{
		Code: code,
		Message: "",
		Data: nil,
	}
}

//返回数据
func (gif *GifeErr) AddData(data interface{}) GifeErr {
	return GifeErr{
		Code: gif.Code,
		Message:  gif.Message,
		Data: data,
	}
}

//添加报错信息
func (gif *GifeErr) AddMessage(message string) GifeErr {
	return GifeErr{
		Code: gif.Code,
		Message: message,
		Data: nil,
	}
}