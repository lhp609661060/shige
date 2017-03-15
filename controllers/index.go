package controllers

import (
	"github.com/astaxie/beego"
	"shige/models"
	_ "fmt"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	c.Data["ShigeDoc"] = models.ShigeData()
	c.TplName = "index.tpl"
}
