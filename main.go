package main

import (
	"github.com/astaxie/beego"
	"shige/filters"
	"shige/models"
	_ "shige/routers"
)

func main() {
	beego.AddFuncMap("replace", filters.Replace)
	s := models.MgoInit()
	defer s.Close()

	beego.Run()
}

