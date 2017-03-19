package admin

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

func (c *IndexController) initTpl() {
	c.Layout = "admin/admin_base.tpl"
	c.TplName = "admin/m.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = ""
	c.LayoutSections["Scripts"] = ""
	c.LayoutSections["Sidebar"] = ""
}

func (c *IndexController) Get() {
	c.initTpl()
	c.Data["Title"] = "诗-后台管理"
}