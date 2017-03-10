package controllers

import (
	"github.com/astaxie/beego"
)

// IndexController operations for Index
type IndexController struct {
	beego.Controller
}

// URLMapping ...
func (c *IndexController) URLMapping() {
	c.Mapping("/", c.Get)
}

func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
