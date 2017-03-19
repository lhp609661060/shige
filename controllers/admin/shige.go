package admin

import (
	"github.com/astaxie/beego"
	"shige/models"
	"gopkg.in/mgo.v2/bson"
)

type ShigeController struct {
	beego.Controller
}

func (c *ShigeController) Get() {

	data := models.ShigeDataByQ(bson.M{})

	c.Data["json"] = map[string]interface{}{
		"status": 1,
		"message": "test",
		"data": data,
	}
	c.ServeJSON()
}