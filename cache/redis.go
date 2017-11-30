package cache

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
)

var (
	//RedisClient 操作客户端
	RedisClient *redis.Pool
)

// RedisConfig 配置
type RedisConfig struct {
	Url string  //地址
	Port string //端口
	Password string //密码
	DB int    //
}

//redis init

func init(){
	RedisClient = newPool()
}
func newPool() *redis.Pool {
	config:=new(RedisConfig)
	config.Url = "127.0.0.1"
	config.Port = "6379"
	config.Password = "mypassword"
	config.DB = 1

	return  &redis.Pool{
		MaxActive:10000,
		MaxIdle:10,
		IdleTimeout:30*time.Second,
		Dial:func()(redis.Conn,error){
			c,err:=redis.Dial("tcp",fmt.Sprintf("%s:%s",config.Url,config.Port),)
			if err != nil{
				 return nil,err
			}
			if _,err:= c.Do("AUTH",config.Password) err!=nil {
				c.Close()
				return nil,err
			}
			if _,err= c.Do("SELECT",config.DB) err!=nil {
				c.Close()
				return nil,err
			}
			return c,nil
		},
	}
}