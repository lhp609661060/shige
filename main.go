package main

import (
	_ "shige/routers"
	"github.com/astaxie/beego"
	"shige/filters"
)

func main() {
	beego.AddFuncMap("replace", filters.Replace)
	beego.Run()
}

