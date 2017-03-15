package main

import (
	_ "shige/routers"
	"github.com/astaxie/beego"
	"shige/filters"
	"shige/cache"
)

func main() {
	beego.AddFuncMap("replace", filters.Replace)
	cache.InitRedisClient()
	cache.CacheRedisClientTest()
	beego.Run()
}

