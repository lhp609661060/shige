package admin

import (
	"github.com/astaxie/beego"
	"shige/models"
	_ "fmt"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) initTpl() {
	c.Layout = "admin/admin_base.tpl"
	c.TplName = "admin/login.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "admin/login_css.tpl"
	c.LayoutSections["Scripts"] = ""
	c.LayoutSections["Sidebar"] = ""
}

func (c *LoginController) Get() {
	c.Data["InputName"] = ""
	c.initTpl()
}

func (c *LoginController) Post() {
	c.initTpl()
	username := c.GetString("username")
	password := c.GetString("password")

	isAuth, inputName, errMsg := models.AuthUser(username, password)
	if isAuth {
		c.Redirect("/admin/m/", 302)
	}else{
		c.Data["InputName"] = inputName
		c.Data["ErrMsg"] = errMsg
	}
}
