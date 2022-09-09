package lib

import "github.com/go-redis/redis"

//加载redis 协程安全的经过测试发现 初始化加载
var Redis *redis.Client = nil
//初始化redis服务配置
func InitRedis(linkStr string) {
	opt, err := redis.ParseURL(linkStr)
	if err != nil {
		panic(err)
	}
	Redis = redis.NewClient(opt)
}


