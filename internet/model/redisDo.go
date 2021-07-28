package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

//往redis存入map类型
func HashSet(key string, mapData map[string]interface{}) (bool, string) {
	//参数验证
	if key == "" || mapData == nil {
		return false,"参数为空"
	}

	if RedisClient == nil {
		fmt.Println("客户端为空")
		return false,"客户端链接断开"
	}else {
		if err := RedisClient.HMSet(key, mapData).Err(); err != nil {
			fmt.Println(err)
			return false, "添加失败"
		} else {
			return true, ""
		}
	}
}

//往redis存入string类型
func StringSet(key string, value int) (bool, string) {
	//参数验证
	if key == "" {
		return false,"参数为空"
	}
	if err := RedisClient.Set(key, value, 0).Err();err!=nil{
		fmt.Println(err)
		return false, "添加失败"
	}else {
		return true,""
	}
}

//从redis中取map所有
func HashGetAll(key string) (interface{},string){
	//参数非空判断
	if key == "" {
		return "","参数为空"
	}
	value, err := RedisClient.HGetAll(key).Result()
	if err == redis.Nil{
		return "","key不存在"
	}else if err!= nil {
		return "","获取失败"
	}
	return value,""
}

//从redis中取string类型值
func StringGet(key string) (string,string) {
	//参数非空判断
	if key == "" {
		return "","参数为空"
	}
	value,err := RedisClient.Get(key).Result()
	if err != nil{
		fmt.Println(err)
	}
	return value,""
}

//从redis的map中取一个属性
func HashGet(key string,field string) (string){
	//参数非空判断
	if key == "" || field == "" {
		return ""
	}
	value := RedisClient.HGet(key,field).Val()
	return value
}

//判断是否含有该key
func HasKey(keyname string) (bool) {
	//参数非空判断
	if keyname == "" {
		return false
	}
	exists,err := RedisClient.Exists(keyname).Result()
	if err != nil {
		fmt.Println(err)
	}
	if exists == 1 {
		return true
	}else {
		return false
	}
}

//获取map中的所有value值
func HasHashKay(keyname string) []string {
	//参数非空判断
	if keyname == "" {
		return nil
	}
	exists,err := RedisClient.HKeys(keyname).Result()
	if err != nil {
		fmt.Println(err)
	}
	return exists
}

//redis中map修改值
func ValueUpdate(key string, field string, newvalue interface{}) (error) {
	_, err := RedisClient.HSet(key, field, newvalue).Result()
	if err!=nil{
		fmt.Println(err)
		return err
	}
	return nil
}

//redis中原子增加操作
func Incr(key string) (int64) {
	result, err := RedisClient.Incr(key).Result()
	if err!=nil{
		fmt.Println(err)
	}
	return result
}

//redis中原子减操作
func Decr(key string) int64 {
	result, err := RedisClient.Decr(key).Result()
	if err!=nil{
		fmt.Println(err)
	}
	return result
}

//redis中删除hash字段
func Hdel(key string, field string) int64 {
	result, err := RedisClient.HDel(key, field).Result()
	if err!=nil{
		fmt.Println(err)
	}
	return result
}
