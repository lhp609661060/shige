package controllers

import (
	"github.com/astaxie/beego"
	"../models"
	"fmt"
)

// IndexController operations for Index
type IndexController struct {
	beego.Controller
}
//
//// URLMapping ...
//func (c *IndexController) URLMapping() {
//	c.Mapping("/", c.Get)
//}

func (c *IndexController) Get() {
	c.Data["Website"] = models.TestString()
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["ShigeDoc"] = models.TestData()
	fmt.Println(models.TestData())
	c.TplName = "index.tpl"
}
