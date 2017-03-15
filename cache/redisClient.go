package cache

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
	"time"
)

var __redisClien cache.Cache
var err error

func GetCache () cache.Cache{
	return __redisClien
}

func InitRedisClient() {
	__redisClien, err = cache.NewCache("redis", `{"conn":"10.57.115.180:6379"}`)
	if err != nil {
		fmt.Println("redis connect error")
		fmt.Println(err)
	}
}

func CacheRedisClientTest()  {
	__redisClien.Put("goredistest", "ces", 60 * time.Second)
	fmt.Println("goredistest")
	data := __redisClien.Get("goredistest")
	fmt.Printf("%T\n", data);
	fmt.Println(data)
}
