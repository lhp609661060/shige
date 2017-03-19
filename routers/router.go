package routers

import (
	"shige/controllers"
	"shige/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})

	ns :=
		beego.NewNamespace("/admin",
			beego.NSRouter("/login", &admin.LoginController{}),
			beego.NSRouter("/m", &admin.IndexController{}),
			beego.NSRouter("/shige", &admin.ShigeController{}),
		)
	//注册 namespace
	beego.AddNamespace(ns)
}
